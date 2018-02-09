package noaa

import (
	"encoding/json"
	"fmt"
	"github.com/buskersguidetotheuniverse.org/net"
	"github.com/buskersguidetotheuniverse.org/types"
	"log"
)

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
