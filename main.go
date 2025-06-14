package main

import (
	"flag"
	"fmt"

	"weather-app/geo"
)

func main() {
	city := flag.String("city", "", "City of the user")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println("Error getting my location:", err)
	}
	fmt.Println(geoData)
}
