package noaa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/buskersguidetotheuniverse.org/types"
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
	body, err := readFromUrl(apiUrl)

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
	body, err := readFromUrl(apiUrl)

	var data types.CurrentConditionsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, err
}


func readFromUrl(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// eventually we need proper error handling, as there are a number of reasons why it's reasonable
		// to expect an error here.
		log.Fatal(err)
	}

	req.Header.Set("UserAgent", "student experiment for reading JSON") // NWS requires a User-Agent
	req.Header.Set("Accept", "*")                                      // NWS requires accept

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()


	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(readErr)
	}

	return body, readErr

}