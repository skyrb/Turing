package main

import (
	"fmt"
	"math/rand"
	"time"
)

// WeatherData represents the structure for holding weather data for a region.
type WeatherData struct {
	Temperature float64
	Humidity    float64
}

// simulateWeatherAPI simulates fetching weather data from an API.
// In reality, you would replace this with actual API calls.
func simulateWeatherAPI(region string) WeatherData {
	// Simulating temperature and humidity data
	temperature := 15.0 + rand.Float64()*(35.0-15.0) // Random temperature between 15 and 35
	humidity := 40.0 + rand.Float64()*(80.0-40.0)    // Random humidity between 40 and 80
	return WeatherData{
		Temperature: temperature,
		Humidity:    humidity,
	}
}

// collectWeatherData collects weather data for the given regions.
func collectWeatherData(regions []string) map[string]WeatherData {
	data := make(map[string]WeatherData)
	for _, region := range regions {
		data[region] = simulateWeatherAPI(region)
	}
	return data
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	regions := []string{"New York", "California", "Texas", "Florida", "Illinois"}

	// Collect weather data for the regions
	weatherData := collectWeatherData(regions)

	// Print the collected data
	for region, data := range weatherData {
		fmt.Printf("Region: %s, Temperature: %.2f°C, Humidity: %.2f%%\n",
			region, data.Temperature, data.Humidity)
	}
}