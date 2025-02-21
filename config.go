package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load variables from .env file, if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Now you can access environment variables as usual
	port := os.Getenv("WEATHER_API_KEY")
	if port == "" {
		port = "8080" // fallback value
	}
	fmt.Println("Server running on port:", port)
}
