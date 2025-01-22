package main

import (
	"log"
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

// CalibrationRule holds the calibration adjustments for a sensor
type CalibrationRule struct {
	TemperatureAdjustment float64
}

// AggregatedData holds the aggregated temperature and humidity data
type AggregatedData struct {
	AverageTemperature float64
	AverageHumidity    float64
	Count              int
}

var (
	sensorCalibration = make(map[string]CalibrationRule)
	dataChannel       = make(chan SensorData, 10)
	aggregationPeriod = 10 * time.Second
)

func main() {
	var wg sync.WaitGroup

	// Register sensors with their calibration rules
	registerSensor("Sensor1", CalibrationRule{TemperatureAdjustment: -1.5})
	registerSensor("Sensor2", CalibrationRule{TemperatureAdjustment: 2.0})

	wg.Add(1)
	go receiveData(&wg)

	wg.Add(1)
	go aggregateData(&wg)

	// Wait for all Goroutines to finish
	wg.Wait()
}

// registerSensor registers a new sensor with its calibration rules
func registerSensor(sensorID string, rule CalibrationRule) {
	sensorCalibration[sensorID] = rule
}

// receiveData simulates receiving real-time data from sensors
func receiveData(wg *sync.WaitGroup) {
	defer wg.Done()

	sensors := []string{"Sensor1", "Sensor2"}
	for i := 0; i < 100; i++ { // Simulating 100 data points
		for _, sensorID := range sensors {
			sensorData := SensorData{
				SensorID:    sensorID,
				Temperature: rand.Float64()*40 - 5, // Random temp between -5 and 35°C
				Humidity:    rand.Float64() * 100,  // Random humidity between 0 and 100%
			}
			calibratedData := calibrateSensorData(sensorData)
			dataChannel <- calibratedData
		}
		time.Sleep(500 * time.Millisecond) // Simulate delay between readings
	}
	close(dataChannel)
}

// calibrateSensorData applies calibration rules to the sensor data
func calibrateSensorData(data SensorData) SensorData {
	if rule, exists := sensorCalibration[data.SensorID]; exists {
		data.Temperature += rule.TemperatureAdjustment
	}
	return data
}

// aggregateData aggregates data every 10 seconds and calculates average temperature and humidity
func aggregateData(wg *sync.WaitGroup) {
	defer wg.Done()

	aggregationTicker := time.NewTicker(aggregationPeriod)
	defer aggregationTicker.Stop()

	aggregatedData := make(map[string]AggregatedData)

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				// Channel is closed, exit the loop
				return
			}
			aggregateSensorData(aggregatedData, data)

		case <-aggregationTicker.C:
			logAggregatedData(aggregatedData)
			// Reset aggregated data for the next period
			aggregatedData = make(map[string]AggregatedData)
		}
	}
}

// aggregateSensorData updates the aggregated data with new sensor data
func aggregateSensorData(aggregatedData map[string]AggregatedData, data SensorData) {
	entry := aggregatedData[data.SensorID]
	entry.AverageTemperature = ((entry.AverageTemperature * float64(entry.Count)) + data.Temperature) / float64(entry.Count+1)
	entry.AverageHumidity = ((entry.AverageHumidity * float64(entry.Count)) + data.Humidity) / float64(entry.Count+1)
	entry.Count++
	aggregatedData[data.SensorID] = entry
}

// logAggregatedData logs the aggregated data for each sensor
func logAggregatedData(aggregatedData map[string]AggregatedData) {
	for sensorID, data := range aggregatedData {
		log.Printf("Aggregated Data for %s: Avg Temperature: %.2f°C, Avg Humidity: %.2f%%, Count: %d",
			sensorID, data.AverageTemperature, data.AverageHumidity, data.Count)
	}
}