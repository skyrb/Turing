package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WeatherData represents the structure for holding weather data for a region.
type WeatherData struct {
	Temperature float64
	Humidity    float64
}

var weatherMap sync.Map // Concurrent map to store weather data

// simulateWeatherAPI simulates fetching weather data from a station.
// In practice, you would replace this with actual data fetching logic.
func simulateWeatherAPI(stationID, region string) WeatherData {
	// Simulating data
	temperature := 20.0 + rand.Float64()*(30.0-20.0) // Random temperature between 20 and 30
	humidity := 30.0 + rand.Float64()*(70.0-30.0)    // Random humidity between 30 and 70
	fmt.Printf("Station %s - Region: %s, Temperature: %.2f°C, Humidity: %.2f%%\n",
		stationID, region, temperature, humidity)
	return WeatherData{
		Temperature: temperature,
		Humidity:    humidity,
	}
}

// updateWeatherData safely updates the weather data for a given region.
func updateWeatherData(region string, data WeatherData) {
	weatherMap.LoadOrStore(region, data)
	weatherMap.Store(region, data)
}

// collectDataFromStations simulates data collection from different stations for the same region.
func collectDataFromStations(region string, numStations int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numStations; i++ {
		stationID := fmt.Sprintf("%s-Station-%d", region, i+1)
		data := simulateWeatherAPI(stationID, region)
		updateWeatherData(region, data)
		time.Sleep(time.Millisecond * 100) // Simulate slight delays between station updates
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	regions := []string{"Region A", "Region B"}
	var wg sync.WaitGroup

	for _, region := range regions {
		wg.Add(1)
		go collectDataFromStations(region, 5, &wg) // Simulating 5 stations per region
	}

	wg.Wait()

	// Print the latest weather data for all regions
	fmt.Println("\nFinal weather data for each region:")
	weatherMap.Range(func(region, data interface{}) bool {
		weather := data.(WeatherData)
		fmt.Printf("Region: %s, Temperature: %.2f°C, Humidity: %.2f%%\n",
			region, weather.Temperature, weather.Humidity)
		return true
	})
}