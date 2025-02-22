// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Optionally, pre-load the cache at startup.
	if err := refreshCache(); err != nil {
		log.Fatalf("Error initializing weather data: %v", err)
	}

	http.HandleFunc("/weather", weatherHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
