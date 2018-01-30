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
const nearestStationsEndpoint = apiBaseUrl + "points/%v/stations"
const defaultStation = "KMSP"

//const apiUrl = 	"http://localhost:8080/stations/local/observations/current"

func NearestStations(latitude string, longitude string, limit int) [] string {
	var stations []string

	for i := 0; i < limit; i++ {
		stations = append(stations, "KMSP") // TODO: fetch the list using the nearestStationsEndpoint
	}

	return stations
}


func CurrentConditions(stationId string) (types.CurrentConditionsResponse, error) {
	var err error = nil

	if len(stationId) == 0 {
		stationId = defaultStation
	}

	apiUrl := fmt.Sprintf(observationsEndpoint, stationId)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
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

	var data types.CurrentConditionsResponse

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(readErr)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, err
}
