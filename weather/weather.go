package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"weather-app/geo"
)

var ErrorInvalidFormat = errors.New("invalid format")

func GetWeather(geoData geo.GeoData, format int) (string, error) {
	if format < 1 || format > 3 {
		return "", ErrorInvalidFormat
	}
	baseURL, err := url.Parse("https://wttr.in/" + geoData.City)
	if err != nil {
		return "", errors.New("error parsing URL")
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return "", errors.New("error in http request")
	}
	if resp.StatusCode != 200 {
		return "", errors.New("failed to get weather")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error reading body of response")
	}
	return string(body), nil
}
