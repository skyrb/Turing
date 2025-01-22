package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// SensorData holds the temperature and humidity readings
type SensorData struct {
	Temperature float64
	Humidity    float64
}

func main() {
	var wg sync.WaitGroup

	// Channel to receive sensor data
	dataChannel := make(chan SensorData, 5)

	// Simulate receiving data
	wg.Add(1)
	go receiveData(&wg, dataChannel)

	// Validate data
	wg.Add(1)
	go validateData(&wg, dataChannel)

	// Wait for all Goroutines to finish
	wg.Wait()
}

// receiveData simulates receiving real-time data from sensors
const maxRetries = 3 // Maximum number of retries for faulty data

func receiveData(wg *sync.WaitGroup, dataChannel chan<- SensorData) {
	defer wg.Done()
	defer close(dataChannel)

	for i := 0; i < 10; i++ {
		var sensorData SensorData
		for retry := 0; retry <= maxRetries; retry++ {
			sensorData = SensorData{
				Temperature: rand.Float64()*40 - 5,
				Humidity:    rand.Float64() * 100,
			}
			// Simulate faulty data (-10°C or 100°C) for testing
			if sensorData.Temperature == -10 || sensorData.Temperature == 100 {
				log.Printf("Received faulty data: %+v (Retry %d)", sensorData, retry+1)
				time.Sleep(1 * time.Second) // Delay before retrying
			} else {
				break // Valid data received, break the retry loop
			}
		}

		if isValidTemp(sensorData.Temperature) {
			fmt.Printf("Received data: %+v\n", sensorData)
			dataChannel <- sensorData
		} else {
			// Alert the user or monitoring system about the faulty sensor
			log.Printf("Faulty sensor detected: %+v (Threshold exceeded and no valid data after %d retries)", sensorData, maxRetries)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// validateData checks if the temperature is within the threshold and logs the result
func validateData(wg *sync.WaitGroup, dataChannel <-chan SensorData) {
	defer wg.Done()

	for sensorData := range dataChannel {
		if isValidTemp(sensorData.Temperature) {
			log.Printf("Valid data: Temperature: %.2f°C, Humidity: %.2f%%", sensorData.Temperature, sensorData.Humidity)
		} else {
			log.Printf("Threshold exceeded: Temperature: %.2f°C is out of range!", sensorData.Temperature)
		}
	}
}

// isValidTemp checks if the temperature is within the valid range (15-35°C)
func isValidTemp(temp float64) bool {
	return temp >= 15 && temp <= 35
}