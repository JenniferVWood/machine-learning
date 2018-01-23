package main

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/weather"
)

// so far, just get the data and spit some of it to console
func main() {
	fmt.Printf("Current Conditions:\n")

	currentConditions := weather.CurrentConditions("KMSP")
	fmt.Println(currentConditions.Props.Station)
	fmt.Println(currentConditions.Props.Timestamp)
	//fmt.Println(currentConditions.Props.RawMessage)
	//fmt.Println(currentConditions.Props.BarometricPressure)
	//fmt.Println(currentConditions.Props.Temperature)
	//fmt.Println(currentConditions.Props.WindSpeed)
	//fmt.Println(currentConditions.Props.WindDirection)
	fmt.Println(currentConditions.Props.PrecipitationLastHour)
	fmt.Println(currentConditions.Props.TextDescription)
	//fmt.Println(currentConditions.Props.PresentWeather)

	key, err := weather.MakeKeyFromNoaaTimeStamp("KMSP", currentConditions.Props.Timestamp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("generated key: %v\n", key)

}
