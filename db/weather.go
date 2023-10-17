package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Timezone            string  `json:"timezone"`
	CurrentWeatherUnits struct {
		Time        string `json:"time"`
		Temperature string `json:"temperature"`
		Windspeed   string `json:"windspeed"`
	} `json:"current_weather_units"`
	CurrentWeather struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

var baseUrl = "https://api.open-meteo.com/v1/forecast?current_weather=true"

// GetWeather retrieves the weather information based on the given longitude and latitude.
//
// Parameters:
// - longitude: the longitude of the location.
// - latitude: the latitude of the location.
//
// Returns:
// - WeatherResponse: the weather information for the specified location.
// - error: an error if there was a problem retrieving the weather information.
func GetWeather(longitude, latitude float64) (WeatherResponse, error) {
	apiUrl := fmt.Sprintf("%s&latitude=%g&longitude=%g", baseUrl, longitude, latitude)
	var weatherResponse WeatherResponse

	// get response unmarshal
	res, err := http.Get(apiUrl)
	if err != nil {
		return weatherResponse, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return weatherResponse, err
	}
	if err := json.Unmarshal(resBody, &weatherResponse); err != nil {
		return weatherResponse, err
	}

	return weatherResponse, nil
}
