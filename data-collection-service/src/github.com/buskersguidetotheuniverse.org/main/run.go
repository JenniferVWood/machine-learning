package main

import (
	"flag"
	"fmt"
	"github.com/Tkanos/gonfig"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/buskersguidetotheuniverse.org/openei"
	"github.com/buskersguidetotheuniverse.org/types"
	"log"
	"os"
	"strconv"
	"sync"
)

// Fetch the weather from a series of NOAA stations and save the results to a local hbase instance.
func main() {

	printWeather := flag.Bool("report", false, "print conditions for each stations to console")

	latitude := flag.String("lat", "", "(optional, but must be used with long) search for stations near latitude")
	longitude := flag.String("long", "", "(optional, but must be used with longitude) search for stations near longitude")
	maxStations := flag.Int("limit", 0, "(optional, but must be used with longitude and latitude) limit number of stations near coordinates")

	flag.Parse()

	stations := flag.Args()

	numStations := len(stations)
	if numStations == 0 && *latitude == "" && *longitude == "" && *maxStations == 0 {
		fmt.Printf("No stations passed in.")
		os.Exit(0)
	}

	log.Printf("%v", os.Args)
	log.Printf("-p: %v\n", *printWeather)
	log.Printf("tail: %v\n", stations)

	configuration := types.Configuration{}
	err := gonfig.GetConf("/home/jennifer/src/java/machine-learning/data-collection-service/config.json", &configuration)
	if err != nil {
		panic(err)
	}

	// todo:  there's got to be a better way to do the lat/long conversion
	lat, err := strconv.ParseFloat(*latitude, 32)
	lon, err := strconv.ParseFloat(*longitude, 32)

	var coords [2]float64
	coords[0] = lat
	coords[1] = lon
	queryLocation := types.Geometry{
		Coordinates: coords,
	}

	var wg sync.WaitGroup

	// first the weather
	if *latitude != "" && *longitude != "" && *maxStations > 0 {
		handleLatLong(&queryLocation, printWeather, &wg)
	}

	if len(stations) > 0 {
		handleStationLiterals(stations, printWeather, &wg)
	}

	// now the energy pricing
	energyClient := openei.NewClient(configuration.Key)
	prices, _ := energyClient.CurrentEnergyPrices(queryLocation)
	err = hbase.SaveEnergyPrices(&prices)

	log.Println("waiting for all threads to return.")
	wg.Wait()
	log.Println("Exiting normally.")
}

func handleStationLiterals(stations []string, printWeather *bool, wg *sync.WaitGroup) {
	for _, station := range stations {
		wg.Add(1)
		s := station
		request := types.ProcessStationRequest{
			PrintWeather: *printWeather,
			StationId:    s,
		}
		go ProcessStation(&request, wg)
	}
}

func handleLatLong(queryLocation *types.Geometry, printWeather *bool, wg *sync.WaitGroup) {

	nearestStations, err := noaa.NearestStations(queryLocation)
	if err != nil {
		log.Fatalf("Error getting stations for point (%v): %v", queryLocation, err)
		os.Exit(1)
	}

	for _, station := range nearestStations.Features {
		wg.Add(1)

		request := types.ProcessStationRequest{
			StationProperties: station.Properties,
			PrintWeather:      *printWeather,
			QueryLocation:     *queryLocation,
		}
		go ProcessStation(&request, wg)
	}

}

// TODO: This is both NOAA and HBase code.  Maybe it belongs somewhere else?
// Download current observations for a given station, and persist the results to our HBase instance.
// For now, we discard information about the actual station.  Most of what we need is embedded in the conditions response.
//
func ProcessStation(request *types.ProcessStationRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	currentConditions, err := noaa.CurrentConditions(request.StationProperties.StationIdentifier)
	if err != nil {
		log.Printf("WARN: couldn't look up weather for %v:  %v", request.StationProperties.StationIdentifier, err)

		// not a fatal error, but obviously we don't want to try to persist the result
		return
	}

	currentConditions.Props.QueryLocation = request.QueryLocation

	//TODO: Use ErrorGroup or whatever instead.
	err = hbase.SaveObservation(&currentConditions)
	if err != nil {
		log.Fatalf("Error saving weather to HBase: %v", err)
		os.Exit(1)
	}

	if request.PrintWeather {
		printCurrentConditions(&currentConditions)
	}

	//log.Println("done processing")

}

func printCurrentConditions(currentConditions *types.CurrentConditionsResponse) {
	fmt.Printf("Current Conditions:\n")
	fmt.Println(currentConditions.Props.Station)
	fmt.Println(currentConditions.Props.Timestamp)
	//fmt.Println(currentConditions.Props.BarometricPressure)
	fmt.Println(currentConditions.Props.Temperature)
	fmt.Println(currentConditions.Props.WindSpeed)
	fmt.Println(currentConditions.Props.WindDirection)
	fmt.Println(currentConditions.Props.TextDescription)

}
