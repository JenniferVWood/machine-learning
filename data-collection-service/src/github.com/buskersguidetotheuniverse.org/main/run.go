package main

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/types"
	"flag"
	"log"
	"os"
	"reflect"
)

// Fetch the weather from a series of NOAA stations and save the results to a local hbase instance.
func main() {

	printWeather := flag.Bool("report", false, "print conditions for each stations to console")
	flag.Parse()

	stations := flag.Args()

	numStations := len(stations)
	if numStations == 0 {
		fmt.Printf("No stations passed in.")
		os.Exit(0)
	}

	log.Printf("%v", os.Args)
	log.Printf("-p: %v\n", *printWeather)
	log.Printf("tail: %v\n", stations)


	var chans [] chan int
	for i := 0; i < numStations; i++ {
		ch := make(chan int)
		chans = append(chans, ch)
		go processStation(stations[i], *printWeather, ch)
	}


	// I wish I really understood how this worked...
	// taken from https://play.golang.org/p/8zwvSk4kjx
	// and https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement
	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	remaining := len(cases)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			// The chosen channel has been closed, so zero out the channel to disable the case
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining -= 1
			continue
		}

		log.Printf("Read from channel %#v and received %s\n", chans[chosen], value.String())
	}
}

func processStation(station string, printWeather bool, ch chan int) {
	currentConditions, err := noaa.CurrentConditions(station)
	if err != nil {
		log.Printf("WARN: couldn't look up weather for %v:  %v", station, err)

		// not a fatal error, but obviously we don't want to try to persist the result
		return
	}

	err = hbase.SaveObservation(&currentConditions)
	if err != nil {
		log.Fatalf("Error saving weather to HBase: %v", err)
		os.Exit(1)
	}

	if printWeather {
		printCurrentConditions(&currentConditions)
	}

	ch <- 1
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
