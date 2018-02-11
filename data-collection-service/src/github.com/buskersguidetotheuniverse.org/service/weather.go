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

type WeatherService struct {
}

func NewWeatherService() WeatherService {
	return WeatherService{}
}

// TODO: do we need to worry about concurrency here?  I don't think so?
// Download current observations for a given station, and persist the results to our HBase instance.
// For now, we discard information about the actual station.  Most of what we need is embedded in the conditions response.
func (ws WeatherService) ProcessStation(request *types.ProcessStationRequest, wg *sync.WaitGroup) {
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
