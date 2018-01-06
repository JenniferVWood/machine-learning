package weather;

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
 *	Get data from NOAA
 *
 * https://api.weather.gov/stations/KMSP/observations/current
 */
//const apiUrl = "https://api.weather.gov/stations/{stationId}/observations/current"
const apiUrl = "https://api.weather.gov/stations/KMSP/observations/current"
//const apiUrl = "https://api.weather.gov/stations?limit=1"
//const apiUrl = 	"http://localhost:8080/stations/local/observations/current"

	func Fetch(stationId string) {
	// for now, ignore the stationId -- we're only interested in kMSP

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("UserAgent", "student experiment for reading JSON") // NWS requires a User-Agent
	req.Header.Set("Accept", "*") // NWS requires accept

	res, _ := client.Do(req)
	defer res.Body.Close()



	var data interface{}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(readErr)
	}

	_ = json.Unmarshal(body, &data)

	fmt.Print(data)

}