package service

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/buskersguidetotheuniverse.org/types"
	"log"
	"os"
	"sync"
)

// It is up to calling code to call WaitFor() on the passed-in sync.WaitGroup.
type WeatherService struct {
	wg *sync.WaitGroup
}

// It is up to calling code to call WaitFor() on the passed-in sync.WaitGroup.
func NewWeatherService(wg *sync.WaitGroup) WeatherService {
	return WeatherService{
		wg: wg,
	}
}

type ProcessStationRequest struct {
	StationProperties types.StationProperties // populated when operating on a lat-long
	StationId         string                  // when stationId is passed in explicitly
	QueryLocation     types.Geometry
	PrintWeather      bool
}

func (ws WeatherService) ProcessStations(stations []string, printWeather bool) {

	for _, station := range stations {
		ws.wg.Add(1)
		s := station
		request := ProcessStationRequest{
			PrintWeather: printWeather,
			StationId:    s,
		}
		go ws.ProcessStation(&request)
	}
}

// TODO: do we need to worry about concurrency here?  I don't think so?
// it all comes down to this function
// Download current observations for a given station, and persist the results to our HBase instance.
func (ws WeatherService) ProcessStation(request *ProcessStationRequest) {
	defer ws.wg.Done()

	currentConditions, err := noaa.CurrentConditions(request.StationProperties.StationIdentifier)
	if err != nil {
		log.Printf("WARN: couldn't look up weather for %v:  %v", request.StationProperties.StationIdentifier, err)

		// not a fatal error, but obviously we don't want to try to persist the result
		return
	}

	currentConditions.Props.QueryLocation = request.QueryLocation
	currentConditions.Props.DistanceFromQueryLoc = GetDistance(&request.QueryLocation, &currentConditions.Geometry)

	//TODO: Use ErrorGroup or whatever instead.
	err = hbase.SaveObservation(&currentConditions)
	if err != nil {
		log.Fatalf("Error saving weather to HBase: %v", err)
		os.Exit(1)
	}

	if request.PrintWeather {
		printCurrentConditions(&currentConditions)
	}
}

// Look up the weather stations near the queryLocation, and get reports for each.
func (ws WeatherService) ProcessLatLong(queryLocation *types.Geometry, printWeather *bool) error {
	nearestStations, err := noaa.NearestStations(queryLocation)

	if err != nil {
		log.Fatalf("Error getting stations for point (%v): %v", queryLocation, err)
		os.Exit(1)
	}

	for _, station := range nearestStations.Features {
		// will be marked done() by ws.ProcessStation()
		ws.wg.Add(1)

		request := ProcessStationRequest{
			StationProperties: station.Properties,
			PrintWeather:      *printWeather,
			QueryLocation:     *queryLocation,
		}
		go ws.ProcessStation(&request)
	}

	return nil
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
