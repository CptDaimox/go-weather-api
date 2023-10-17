package main

import (
	"net/http"

	controller "github.com/CptDaimox/go-weather-api/controllers"
)

// the entry point of the program.
//
// It sets up the "/weather" route to the WeatherRouter handler function
// and starts the HTTP server listening on port 3000.
func main() {
	http.HandleFunc("/weather", controller.WeatherRouter)
	http.ListenAndServe(":3000", nil)
}
