package main

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/noaa"
	"github.com/buskersguidetotheuniverse.org/persist"
)

// so far, just get the data and spit some of it to console
func main() {
	fmt.Printf("Current Conditions:\n")

	currentConditions := noaa.CurrentConditions("KMSP")
	fmt.Println(currentConditions.Props.Station)
	fmt.Println(currentConditions.Props.Timestamp)
	//fmt.Println(currentConditions.Props.RawMessage)
	//fmt.Println(currentConditions.Props.BarometricPressure)
	fmt.Println(currentConditions.Props.Temperature)
	//fmt.Println(currentConditions.Props.WindSpeed)
	//fmt.Println(currentConditions.Props.WindDirection)
	//fmt.Println(currentConditions.Props.PrecipitationLastHour)
	fmt.Println(currentConditions.Props.TextDescription)
	//fmt.Println(currentConditions.Props.PresentWeather)

	persist.SaveObservation(&currentConditions)

}
