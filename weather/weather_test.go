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
	got, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Assert
	if !strings.Contains(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestGetWeatherWrongFormat(t *testing.T) {
	// Arrange
	expected := "London"
	geoData := geo.GeoData{City: expected}
	format := 100

	// Act
	_, err := weather.GetWeather(geoData, format)

	// Assert
	if err != weather.ErrorInvalidFormat {
		t.Errorf("Expected %v, got %v", weather.ErrorInvalidFormat, err)
	}
}
