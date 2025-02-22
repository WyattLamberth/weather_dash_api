// weather.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	cache         []byte            // Cached API response.
	cacheMutex    sync.RWMutex      // Mutex for safe concurrent access.
	lastFetchTime time.Time         // Time when the cache was last updated.
	cacheTTL      = 5 * time.Minute // Refresh cache every 5 minutes.
)

// apiCall makes a GET request to the given URL and returns the response body.
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

// refreshCache fetches the latest weather data and updates the cache.
func refreshCache() error {
	weatherAPIKey, latitude, longitude, err := getEnvVars()
	if err != nil {
		return err
	}

	// Build URL with minutely data excluded.
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=minutely&appid=%s", latitude, longitude, weatherAPIKey)
	response, err := apiCall(url)
	if err != nil {
		return err
	}

	cacheMutex.Lock()
	cache = response
	lastFetchTime = time.Now()
	cacheMutex.Unlock()
	return nil
}

// weatherHandler serves the weather data. It refreshes the cache if it's older than cacheTTL.
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	cacheMutex.RLock()
	cachedData := cache
	cachedTime := lastFetchTime
	cacheMutex.RUnlock()

	// Refresh cache if needed.
	if cachedData == nil || time.Since(cachedTime) > cacheTTL {
		if err := refreshCache(); err != nil {
			http.Error(w, fmt.Sprintf("Error refreshing weather data: %v", err), http.StatusInternalServerError)
			return
		}
		cacheMutex.RLock()
		cachedData = cache
		cacheMutex.RUnlock()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cachedData)
}
