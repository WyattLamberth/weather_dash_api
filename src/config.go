// config.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

// getEnvVars retrieves the WEATHER_API_KEY, LATITUDE, and LONGITUDE from environment variables,
// parsing the latitude and longitude as float64.
func getEnvVars() (string, float64, float64, error) {
	weatherAPIKey := os.Getenv("WEATHER_API_KEY")

	latStr := os.Getenv("LATITUDE")
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("error parsing LATITUDE: %v", err)
	}

	longStr := os.Getenv("LONGITUDE")
	longitude, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("error parsing LONGITUDE: %v", err)
	}

	return weatherAPIKey, latitude, longitude, nil
}
