package main

import (
	"flag"
	"fmt"

	"weather-app/geo"
	"weather-app/weather"
)

func main() {
	city := flag.String("city", "", "City of the user")
	format := flag.Int("format", 1, "Format of the weather")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println("Error getting my location:", err)
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
