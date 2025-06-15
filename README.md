# Weather App

A simple Go command-line application that fetches weather information for any city or your current location using geolocation services.

## Features

- 🌍 **Automatic Location Detection**: Get weather for your current location using IP geolocation
- 🏙️ **City-based Weather**: Specify any city to get its weather information
- 📊 **Multiple Output Formats**: Choose from 3 different weather display formats
- ✅ **City Validation**: Validates city names before fetching weather data

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd weather-app
   ```

2. **Build the application:**
   ```bash
   go build -o weather-app
   ```

## Usage

The application supports the following command-line flags:

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-city` | string | `""` | Specify the city name to get weather for |
| `-format` | int | `1` | Weather display format (1, 2, or 3) |

### Examples

#### Get weather for your current location (default):
```bash
./weather-app
```

#### Get weather for a specific city:
```bash
./weather-app -city="New York"
./weather-app -city="London"
./weather-app -city="Tokyo"
# Or with space separation:
./weather-app -city "New York"
```

#### Use different weather formats:
```bash
# Format 1 (default) - Basic weather info
./weather-app -city="Paris" -format=1

# Format 2 - Extended weather info
./weather-app -city="Berlin" -format=2

# Format 3 - Detailed weather info
./weather-app -city="Sydney" -format=3
```

#### Combine flags:
```bash
./weather-app -city="Moscow" -format=2
```

### Weather Formats

The application supports 3 different weather display formats:

- **Format 1**: Basic weather information (temperature, conditions)
- **Format 2**: Extended weather information (includes humidity, wind)
- **Format 3**: Detailed weather information (comprehensive weather data)

## How It Works

1. **Location Detection**: 
   - If no city is specified, the app uses IP geolocation to detect your current location
   - If a city is provided, it validates the city name using a population API

2. **Weather Fetching**: 
   - Uses the wttr.in service to fetch weather data
   - Supports multiple output formats for different levels of detail

## API Dependencies

This application relies on the following external services:

- **wttr.in**: Weather data service
- **ipapi.co**: IP geolocation service
- **countriesnow.space**: City validation service

## Error Handling

The application handles various error scenarios:

- Invalid city names
- Network connectivity issues
- Invalid format numbers (must be 1, 2, or 3)
- API service unavailability

## Requirements

- Go 1.24.1 or later
- Internet connection for API calls

## Project Structure

```
weather-app/
├── main.go           # Entry point with flag parsing
├── go.mod           # Go module definition
├── weather/         # Weather fetching logic
│   ├── weather.go
│   └── weather_test.go
└── geo/             # Geolocation and city validation
    ├── geo.go
    └── geo_test.go
```

## Testing

Run the test suite:
```bash
go test ./...
```

## License

This project is open source and available under the MIT License. 
