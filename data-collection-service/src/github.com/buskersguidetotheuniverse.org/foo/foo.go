package main

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/weather"
)

func main() {
	fmt.Printf("Current Conditions:\n")

	currentConditions := weather.CurrentConditions("KMSP")
	fmt.Println(currentConditions.Props.Station)
	fmt.Println(currentConditions.Props.BarometricPressure)
	fmt.Println(currentConditions.Props.Temperature)
	fmt.Println(currentConditions.Props.WindSpeed)
	fmt.Println(currentConditions.Props.WindDirection)
	fmt.Println(currentConditions.Props.PrecipitationLastHour)
	fmt.Println(currentConditions.Props.TextDescription)
	fmt.Println(currentConditions.Props.PresentWeather)


}
