package geo_test

import (
	"testing"
	"weather-app/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "Bishkek"
	expected := geo.GeoData{
		City: "Bishkek",
	}

	// Act
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Error getting city", err)
	}

	// Assert
	if got.City != expected.City {
		t.Errorf("Expected city %v, got %v", expected.City, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Londondandan"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Expected error %v, got %v", geo.ErrorNoCity, err)
	}
}
