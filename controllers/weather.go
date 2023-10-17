package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	service "github.com/CptDaimox/go-weather-api/pkg"
)

type WeatherGetRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// WeatherRouter is the function that handles the routing for weather requests.
func WeatherRouter(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		get(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// extracts the latitude and longitude from the request and passes it to the GetWeather function.
// returns the weather data as a json response.
func get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if len(queryParams) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert the latitude and longitude strings to float
	latitude, err := strconv.ParseFloat(queryParams.Get("lat"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(queryParams.Get("lon"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	weather, err := service.GetWeather(latitude, longitude)
	if err != nil {
		return
	}

	// With request Body
	// reqBody, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	return
	// }
	// var weatherRequest WeatherGetRequest
	// if err := json.Unmarshal(reqBody, &weatherRequest); err != nil {
	// 	return
	// }

	// weather, err := service.GetWeather(weatherRequest.Latitude, weatherRequest.Longitude)
	// if err != nil {
	// 	return
	// }

	json, err := json.Marshal(weather)
	if err != nil {
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
