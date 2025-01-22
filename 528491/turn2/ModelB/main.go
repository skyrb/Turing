package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WeatherData struct {
	Temperature float64
	Humidity    float64
}

func simulateWeatherAPI(region string) WeatherData {
	temperature := 15.0 + rand.Float64()*(35.0-15.0)
	humidity := 40.0 + rand.Float64()*(80.0-40.0)
	return WeatherData{
		Temperature: temperature,
		Humidity:    humidity,
	}
}

func updateWeatherData(regions []string, wg *sync.WaitGroup, weatherDataMap *sync.Map) {
	defer wg.Done()
	for _, region := range regions {
		data := simulateWeatherAPI(region)
		weatherDataMap.Store(region, data)
	}
}

func printWeatherData(weatherDataMap *sync.Map) {
	weatherDataMap.Range(func(key, value interface{}) bool {
		region, _ := key.(string)
		data, _ := value.(WeatherData)
		fmt.Printf("Region: %s, Temperature: %.2f°C, Humidity: %.2f%%\n",
			region, data.Temperature, data.Humidity)
		return true // Continue the range loop
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	regions := []string{"New York", "California", "Texas", "Florida", "Illinois"}

	var weatherDataMap sync.Map

	// Create a WaitGroup to synchronize the update goroutines
	var wg sync.WaitGroup

	wg.Add(1)
	go updateWeatherData(regions, &wg, &weatherDataMap)

	// Wait for all updates to complete
	wg.Wait()

	fmt.Println("\nInitial Weather Data:")
	printWeatherData(&weatherDataMap)

	// Periodic updates (simulated)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("\nUpdating Weather Data...")
		wg.Add(1)
		go updateWeatherData(regions, &wg, &weatherDataMap)
		wg.Wait()

		fmt.Println("\nUpdated Weather Data:")
		printWeatherData(&weatherDataMap)
	}
}