package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define struct to represent sensor data
type SensorData struct {
	Temperature float64
	Humidity    float64
}

// Generate random sensor data
func generateSensorData() SensorData {
	temp := rand.Float64()*20 + 15
	humidity := rand.Float64()*100
	return SensorData{Temperature: temp, Humidity: humidity}
}

// Function to read sensor data and return channel
func readSensorData(wg *sync.WaitGroup) <-chan SensorData {
	dataChan := make(chan SensorData)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			data := generateSensorData()
			dataChan <- data
			time.Sleep(2 * time.Second) // Simulate sensor reading interval
		}
	}()
	return dataChan
}

// Function to validate sensor data and return channel
func validateSensorData(wg *sync.WaitGroup, in <-chan SensorData) <-chan SensorData {
	validDataChan := make(chan SensorData)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range in {
			if data.Temperature >= 15 && data.Temperature <= 35 {
				validDataChan <- data
			}
		}
	}()
	return validDataChan
}

// Function to log sensor data
func logSensorData(wg *sync.WaitGroup, in <-chan SensorData) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range in {
			fmt.Printf("Temperature: %.2f°C, Humidity: %.2f%%\n", data.Temperature, data.Humidity)
		}
	}()
}

func main() {
	var wg sync.WaitGroup

	// Start reading sensor data
	dataChan := readSensorData(&wg)

	// Validate the sensor data
	validDataChan := validateSensorData(&wg, dataChan)

	// Log valid sensor data
	logSensorData(&wg, validDataChan)

	// Wait for all goroutines to finish
	wg.Wait()
}