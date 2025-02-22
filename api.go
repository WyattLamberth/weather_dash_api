package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // autoload the .env file
)

// apiCall makes a GET request to the given URL and returns the response body as a slice of bytes, or an error if something goes wrong.
func apiCall(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// getEnvVars retrieves the WEATHER_API_KEY, LATITUDE, and LONGITUDE from environment variables,
// parses the latitude and longitude as float64, and returns them. If any parsing error occurs, an error is returned.
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

// printPrettyJSON pretty-prints raw JSON bytes with indentation.
func printPrettyJSON(data []byte) {
	var prettyJSON bytes.Buffer
	// Indent the JSON with two spaces per level.
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		fmt.Println("Error formatting JSON:", err)
		fmt.Println("Raw response:", string(data))
		return
	}
	fmt.Println(prettyJSON.String())
}

func main() {
	// Initialize environment variables.
	weatherAPIKey, latitude, longitude, err := getEnvVars()
	if err != nil {
		fmt.Println("Error retrieving environment variables:", err)
		return
	}

	// Debug: print the env variables.
	fmt.Println("weatherAPIKey:", weatherAPIKey)
	fmt.Println("latitude:", latitude)
	fmt.Println("longitude:", longitude)

	// Construct the URL with the correct formatting for float values.
	// For now, we're hardcoding this; in the future, these might be environment variables or toggles from the front end.
	excludeMinutes := true
	var url string

	if excludeMinutes {
		url = fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=minutely&appid=%s", latitude, longitude, weatherAPIKey)
	} else {
		url = fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&appid=%s", latitude, longitude, weatherAPIKey)
	}

	response, err := apiCall(url)
	if err != nil {
		fmt.Println("Error making API call:", err)
		return
	}

	// Pretty-print the JSON response.
	printPrettyJSON(response)
}
