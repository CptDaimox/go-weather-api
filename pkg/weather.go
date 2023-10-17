package service

import (
	database "github.com/CptDaimox/go-weather-api/db"
)

// Retrieves the weather information based on the given longitude and latitude by calling the DB Function
//
// Parameters:
// - longitude: the longitude value of the location.
// - latitude: the latitude value of the location.
//
// Returns:
// - weather: the weather response struct containing the weather information.
// - err: an error if there was a problem retrieving the weather information.
func GetWeather(longitude, latitude float64) (database.WeatherResponse, error) {
	weather, err := database.GetWeather(longitude, latitude)
	if err != nil {
		return database.WeatherResponse{}, err
	}

	return weather, nil
}
