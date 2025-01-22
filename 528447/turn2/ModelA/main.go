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

// Constants for retry mechanism
const (
	maxRetries     = 3
	retryDelay     = time.Second
	validTempMin   = 15.0
	validTempMax   = 35.0
	totalDataPoints = 10
)

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
func receiveData(wg *sync.WaitGroup, dataChannel chan<- SensorData) {
	defer wg.Done()
	defer close(dataChannel) // Close the channel when done

	for i := 0; i < totalDataPoints; i++ { // Simulating 10 data points
		sensorData := SensorData{
			Temperature: rand.Float64()*110 - 10, // Random temp between -10 and 100°C
			Humidity:    rand.Float64() * 100,  // Random humidity between 0 and 100%
		}
		fmt.Printf("Received data: %+v\n", sensorData)
		dataChannel <- sensorData
		time.Sleep(500 * time.Millisecond) // Simulate delay between readings
	}
}

// validateData checks if the temperature is within the threshold and logs the result
func validateData(wg *sync.WaitGroup, dataChannel <-chan SensorData) {
	defer wg.Done()

	for sensorData := range dataChannel {
		if isValidTemp(sensorData.Temperature) {
			log.Printf("Valid data: Temperature: %.2f°C, Humidity: %.2f%%", sensorData.Temperature, sensorData.Humidity)
		} else {
			log.Printf("Invalid data detected: Temperature: %.2f°C", sensorData.Temperature)
			handleAnomaly(sensorData)
		}
	}
}

// isValidTemp checks if the temperature is within the valid range (15-35°C)
func isValidTemp(temp float64) bool {
	return temp >= validTempMin && temp <= validTempMax
}

// handleAnomaly implements retry mechanism and alerts if data is out of range
func handleAnomaly(sensorData SensorData) {
	for retry := 1; retry <= maxRetries; retry++ {
		time.Sleep(retryDelay)
		newData := reFetchData()
		if isValidTemp(newData.Temperature) {
			log.Printf("Recovered data after retry %d: Temperature: %.2f°C, Humidity: %.2f%%", retry, newData.Temperature, newData.Humidity)
			return
		}
		log.Printf("Retry %d failed: Temperature: %.2f°C", retry, newData.Temperature)
	}

	// If all retries failed, send an alert
	sendAlert(sensorData)
}

// reFetchData simulates re-fetching the data from the sensor
func reFetchData() SensorData {
	return SensorData{
		Temperature: rand.Float64()*110 - 10, // New random temp between -10 and 100°C
		Humidity:    rand.Float64() * 100,
	}
}

// sendAlert is called when retries fail and manual intervention is needed
func sendAlert(sensorData SensorData) {
	log.Printf("ALERT! Manual intervention required for Temperature: %.2f°C", sensorData.Temperature)
}