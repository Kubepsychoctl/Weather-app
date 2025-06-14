package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"weather-app/geo"
)

func GetWeather(geoData geo.GeoData, format int) string {
	baseURL, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error getting weather:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Failed to get weather")
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading weather data:", err)
		return ""
	}
	return string(body)
}
