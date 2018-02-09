package net

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadFromUrl(url string) ([]byte, error) {
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
