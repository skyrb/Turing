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
func receiveData(wg *sync.WaitGroup, dataChannel chan<- SensorData) {
	defer wg.Done()
	defer close(dataChannel) // Close the channel when done

	for i := 0; i < 10; i++ { // Simulating 10 data points
		sensorData := SensorData{
			Temperature: rand.Float64()*40 - 5, // Random temp between -5 and 35°C
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
			log.Printf("Threshold exceeded: Temperature: %.2f°C is out of range!", sensorData.Temperature)
		}
	}
}

// isValidTemp checks if the temperature is within the valid range (15-35°C)
func isValidTemp(temp float64) bool {
	return temp >= 15 && temp <= 35
}