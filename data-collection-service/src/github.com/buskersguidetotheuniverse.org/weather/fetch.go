package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
 *	Get data from NOAA
 *
 * default: https://api.weather.gov/stations/KMSP/observations/current
 */
const apiBaseUrl = "https://api.weather.gov/stations/%v/observations/current"
const defaultStation = "KMSP"

//const apiUrl = 	"http://localhost:8080/stations/local/observations/current"

func CurrentConditions(stationId string) CurrentConditionsResponse {

	if len(stationId) == 0 {
		stationId = defaultStation
	}

	apiUrl := fmt.Sprintf(apiBaseUrl, stationId)

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

	var data CurrentConditionsResponse

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(readErr)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))
	return data
}
