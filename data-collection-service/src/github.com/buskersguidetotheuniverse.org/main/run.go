package main

import (
	"flag"
	"fmt"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/buskersguidetotheuniverse.org/types"
	"log"
	"os"
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

	if *latitude != "" && *longitude != "" && *maxStations > 0 {
		nearestStations, err := noaa.NearestStations(*latitude, *longitude)
		if err != nil {
			log.Fatalf("Error getting stations for point (%v, %v): %v", *latitude, *longitude, err)
			os.Exit(1)
		}
		stationIds := noaa.ExtractIdsFromStationsResponse(&nearestStations, *maxStations)
		stations = append(stations, stationIds...)
	}

	numStations := len(stations)
	if numStations == 0 {
		fmt.Printf("No stations passed in.")
		os.Exit(0)
	}

	log.Printf("%v", os.Args)
	log.Printf("-p: %v\n", *printWeather)
	log.Printf("tail: %v\n", stations)

	var wg sync.WaitGroup
	for _, station := range stations {
		s := station
		wg.Add(1)
		go processStation(s, *printWeather, &wg)
	}

	log.Println("waiting for all threads to return.")
	wg.Wait()
	log.Println("Exiting normally.")
}

func processStation(station string, printWeather bool, wg *sync.WaitGroup) {
	//log.Println("processing..." + station)
	defer wg.Done()

	currentConditions, err := noaa.CurrentConditions(station)
	if err != nil {
		log.Printf("WARN: couldn't look up weather for %v:  %v", station, err)

		// not a fatal error, but obviously we don't want to try to persist the result
		return
	}

	//TODO: Use ErrorGroup or whatever instead.
	err = hbase.SaveObservation(&currentConditions)
	if err != nil {
		log.Fatalf("Error saving weather to HBase: %v", err)
		os.Exit(1)
	}

	if printWeather {
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
