package main

import (
	"flag"
	"fmt"
	"github.com/Tkanos/gonfig"
	"github.com/buskersguidetotheuniverse.org/hbase"
	"github.com/buskersguidetotheuniverse.org/openei"
	"github.com/buskersguidetotheuniverse.org/service"
	"github.com/buskersguidetotheuniverse.org/types"
	"log"
	"os"
	"strconv"
	"sync"
)

// Fetch the weather from a series of NOAA stations and save the results to a local hbase instance.
func main() {
	conf, err := config()
	if err != nil {
		panic(err)
	}

	log.Printf("config: %v", *conf)

	var wg sync.WaitGroup

	// first the weather
	ws := service.NewWeatherService(&wg)

	if conf.IncludesCoords && conf.MaxStations > 0 {
		err = ws.ProcessLatLong(&conf.QueryLocation, &conf.PrintWeather)
	}

	if len(conf.StationIds) > 0 {
		ws.ProcessStations(conf.StationIds, conf.PrintWeather)
	}

	// now the energy pricing
	energyClient := openei.NewClient(conf.Key)
	prices, _ := energyClient.CurrentEnergyPrices(conf.QueryLocation)
	err = hbase.SaveEnergyPrices(&prices)

	log.Println("waiting for all threads to return.")
	wg.Wait()
	log.Println("Exiting normally.")
}

func config() (*types.Configuration, error) {
	config := types.Configuration{
		IncludesCoords: false,
	}

	err := gonfig.GetConf("/home/jennifer/src/java/machine-learning/data-collection-service/config.json", &config)
	if err != nil {
		panic(err)
	}

	config.PrintWeather = *flag.Bool("report", false, "print conditions for each stations to console")

	latitude := flag.String("lat", "", "(optional, but must be used with long) search for stations near latitude")
	longitude := flag.String("long", "", "(optional, but must be used with longitude) search for stations near longitude")

	config.MaxStations = *flag.Int("limit", 0, "(optional, but must be used with longitude and latitude) limit number of stations near coordinates")

	flag.Parse()

	// args can be tailed with a list of stationIDs e.g. KMSP KUOW...
	stations := flag.Args()

	config.IncludesCoords = err == nil && (*latitude != "" || *longitude != "")

	if config.IncludesCoords {
		lat, _ := strconv.ParseFloat(*latitude, 32)
		lon, _ := strconv.ParseFloat(*longitude, 32)

		var coords [2]float64
		coords[0] = lat
		coords[1] = lon
		config.QueryLocation = types.Geometry{
			Coordinates: coords,
		}
	}

	numStations := len(stations)
	if numStations == 0 && !config.IncludesCoords {
		fmt.Printf("No stations passed in.")
		os.Exit(0)
	}

	log.Printf("%v", os.Args)
	log.Printf("-p: %v\n", config.PrintWeather)
	log.Printf("tail: %v\n", stations)

	return &config, err
}
