package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

// SensorData holds the temperature and humidity readings
type SensorData struct {
	SensorID    string
	Temperature float64
	Humidity    float64
}

// SensorMetadata holds metadata about a sensor
type SensorMetadata struct {
	ID             string
	CalibrationOffset float64 // Calibration offset for temperature
}

// GreenhouseData holds aggregated data for a greenhouse
type GreenhouseData struct {
	ID             string
	TemperatureSum float64
	TemperatureCount int
	HumiditySum     float64
	HumidityCount   int
	AverageTemperature float64
	AverageHumidity   float64
}

const (
	maxRetries     = 3
	retryDelay     = time.Second
	validTempMin   = 15.0
	validTempMax   = 35.0
	totalDataPoints = 100
	aggregationPeriod = 10 * time.Second
)

var (
	sensorsMetadata = map[string]SensorMetadata{
		"Sensor1": {ID: "Sensor1", CalibrationOffset: 2.0},
		"Sensor2": {ID: "Sensor2", CalibrationOffset: -1.5},
		// Add more sensor metadata as needed
	}
)

func main() {
	var wg sync.WaitGroup
	var greenhouseData sync.Map

	// Channel to receive sensor data
	dataChannel := make(chan SensorData, 5)

	// Channel to update greenhouse data
	updateChannel := make(chan GreenhouseData)

	// Simulate receiving data
	wg.Add(1)
	go receiveData(&wg, dataChannel)

	// Aggregate and analyze data
	wg.Add(1)
	go aggregateData(&wg, dataChannel, updateChannel)

	// Display aggregated data
	wg.Add(1)
	go displayAggregatedData(&wg, updateChannel)

	// Wait for all Goroutines to finish
	wg.Wait()
}

// receiveData simulates receiving real-time data from sensors
func receiveData(wg *sync.WaitGroup, dataChannel chan<- SensorData) {
	defer wg.Done()
	defer close(dataChannel) // Close the channel when done

	for i := 0; i < totalDataPoints; i++ { // Simulating 100 data points
		for sensorID, metadata := range sensorsMetadata {
			sensorData := SensorData{
				SensorID:    sensorID,
				Temperature: rand.Float64()*110 - 10 + metadata.CalibrationOffset, // Random temp between -10 and 100°C with calibration
				Humidity:    rand.Float64() * 100,  // Random humidity between 0 and 100%
			}
			fmt.Printf("Received data: %+v\n", sensorData)
			dataChannel <- sensorData
		}
		time.Sleep(500 * time.Millisecond) // Simulate delay between readings
	}
}

// aggregateData aggregates data for each greenhouse and calculates averages
func aggregateData(wg *sync.WaitGroup, dataChannel <-chan SensorData, updateChannel chan<- GreenhouseData) {
	defer wg.Done()
	ticker := time.NewTicker(aggregationPeriod)
	defer ticker.Stop()

	for {
		select {
		case sensorData := <-dataChannel: