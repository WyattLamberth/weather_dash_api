# Weather Dashboard API

## Overview
A Go-based REST API that provides cached weather data from OpenWeatherMap API with Docker support.
To be used with the Weather Dashboard: {insert link to repo here later}

## 🚀 Features
- Real-time weather data from OpenWeatherMap
- In-memory caching with 5-minute TTL
- Thread-safe concurrent access
- Containerized deployment
- Environment-based configuration

## 📋 Prerequisites
- Go 1.24 or higher
- Podman or Docker (optional)
- OpenWeatherMap API key

## ⚙️ Configuration
Create a `.env` file in the root directory:

```env
WEATHER_API_KEY=your_api_key_here
LATITUDE=your_latitude
LONGITUDE=your_longitude
```

## 🏃‍♂️ Running Locally

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
cd src
go run .
```

Server starts at `http://localhost:8080`

## 🐳 Container Usage (Recommended)

Build the container using Podman:
```bash
podman build -t weather_dash_api .
```

Using Docker:
```bash
docker build -t weather_dash_api -f Containerfile .
```

Run with environment variables on Podman:
```bash
podman run -it --rm -p 8080:8080 --env-file .env weather_dash_api
```

on Docker:
```bash
docker run -it --rm -p 8080:8080 --env-file .env weather_dash_api
```

## 🔌 API Endpoints

### GET /weather

Returns cached weather data in JSON format including:

#### Response Format
```json
{
    "lat": float,           // Latitude of location
    "lon": float,           // Longitude of location
    "timezone": string,     // Location timezone
    "current": {
        "temp": float,      // Current temperature
        "feels_like": float,// Feels like temperature
        "humidity": int,    // Humidity percentage
        "wind_speed": float,// Wind speed
        "weather": [{       // Weather conditions
            "main": string, // Weather main description
            "description": string // Detailed description
        }]
    },
    "hourly": [{           // 48 hour forecast
        "dt": timestamp,   // Time of forecasted data
        "temp": float,     // Temperature
        "weather": [...]   // Weather conditions
    }],
    "daily": [{           // 7 day forecast
        "dt": timestamp,  // Date of forecasted data
        "temp": {
            "day": float, // Day temperature
            "min": float, // Min temperature
            "max": float  // Max temperature
        },
        "weather": [...], // Weather conditions
        "summary": string // Day summary
    }],
    "alerts": [{          // Weather alerts (if any)
        "event": string,  // Alert type
        "description": string, // Alert description
        "start": timestamp,    // Alert start time
        "end": timestamp      // Alert end time
    }]
}
```

All temperature values are in Kelvin. Wind speed is in meters/second.

[For more detailed information on the API documentation](https://openweathermap.org/api/one-call-3)

## 📁 Project Structure
```plaintext
.
├── .env                # Environment configuration
├── Containerfile      # Docker configuration
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── src/
    ├── config.go      # Environment configuration
    ├── main.go        # Application entry point
    └── weather.go     # Weather service logic
```

## 🛠️ Development

### Building From Source
```bash
cd src
go build -o api .
```

### Running Tests (Work in Progress)
```bash
go test ./...
```

## 📝 License
[MIT License](LICENSE)

## 👥 Contributing
1. Fork the repository
2. Create a feature branch
3. Commit changes
4. Push to the branch
5. Open a Pull Request

## ⚠️ Important Notes
- Weather data is cached for 5 minutes
- API key must be kept secret
- Rate limiting applies to OpenWeatherMap API free tier: up to 1000 calls per day