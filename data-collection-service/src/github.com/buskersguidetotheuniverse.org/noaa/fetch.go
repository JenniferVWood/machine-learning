package noaa

import (
	"encoding/json"
	"fmt"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/net"
	"github.com/buskersguidetotheuniverse.org/types"
	"log"
	"os"
	"sync"
)

// TODO: Wrap all this in a client struct?  Construct a Config object from the command-line args?

/*
 *	Get data from NOAA
 *
 * default: https://api.weather.gov/stations/KMSP/observations/current
 */
const apiBaseUrl = "https://api.weather.gov/"
const observationsEndpoint = apiBaseUrl + "stations/%v/observations/current"
const nearestStationsEndpoint = apiBaseUrl + "points/%v,%v/stations"
const defaultStation = "KMSP"

//const apiUrl = 	"http://localhost:8080/stations/local/observations/current"
//https://api.weather.gov/points/44.9778,-93.2650/stations

func NearestStations(latitude string, longitude string) (types.StationsResponse, error) {
	log.Printf("fetching stations near %v, %v", latitude, longitude)
	// it would make some sense to use the Geometry type here, but it doesn't give us what we really need, which
	// is a guarantee of parameter order.
	apiUrl := fmt.Sprintf(nearestStationsEndpoint, latitude, longitude)
	body, err := net.ReadFromUrl(apiUrl)

	var stations types.StationsResponse
	err = json.Unmarshal(body, &stations)
	if err != nil {
		log.Fatal(err)
	}

	return stations, err
}

func CurrentConditions(stationId string) (types.CurrentConditionsResponse, error) {
	var err error = nil

	if len(stationId) == 0 {
		stationId = defaultStation
	}

	apiUrl := fmt.Sprintf(observationsEndpoint, stationId)
	body, err := net.ReadFromUrl(apiUrl)

	var data types.CurrentConditionsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, err
}

// TODO: This is both NOAA and HBase code.  Maybe it belongs somewhere else?
// Download current observations for a given station, and persist the results to our HBase instance.
// For now, we discard information about the actual station.  Most of what we need is embedded in the conditions response.
//
func ProcessStation(request *types.ProcessStationRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	currentConditions, err := CurrentConditions(request.StationProperties.StationIdentifier)
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
