package weather_test

import (
	"strings"
	"testing"

	"weather-app/geo"
	"weather-app/weather"
)

func TestGetWeather(t *testing.T) {
	// Arrange
	expected := "Bishkek"
	geoData := geo.GeoData{City: expected}
	format := 3

	// Act
	got := weather.GetWeather(geoData, format)

	// Assert
	if !strings.Contains(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
