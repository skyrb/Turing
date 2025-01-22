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

// TemperatureRingBuffer holds the temperature data for the past hour using a ring buffer.
type TemperatureRingBuffer struct {
	buffer   []float64
	size     int
	nextIndex int
}

// NewTemperatureRingBuffer creates a new TemperatureRingBuffer of the specified size.
func NewTemperatureRingBuffer(size int) *TemperatureRingBuffer {
	return &TemperatureRingBuffer{
		buffer:   make([]float64, size),
		size:     size,
		nextIndex: 0,
	}
}

// Add adds a new temperature to the ring buffer.
func (rb *TemperatureRingBuffer) Add(temperature float64) {
	rb.buffer[rb.nextIndex] = temperature
	rb.nextIndex = (rb.nextIndex + 1) % rb.size
}

// GetHighest returns the highest temperature in the ring buffer.
func (rb *TemperatureRingBuffer) GetHighest() float64 {
	highest := rb.buffer[0]
	for _, temp := range rb.buffer[1:] {
		if temp > highest {
			highest = temp
		}
	}
	return highest
}

// GetLowest returns the lowest temperature in the ring buffer.
func (rb *TemperatureRingBuffer) GetLowest() float64 {
	lowest := rb.buffer[0]
	for _, temp := range rb.buffer[1:] {
		if temp < lowest {
			lowest = temp
		}
	}
	return lowest
}

// RegionData holds the weather data for a region, including a ring buffer for temperature data and highest/lowest temperatures.
type RegionData struct {
	TemperatureBuffer *TemperatureRingBuffer
	HighestTemperature float64
	LowestTemperature  float64
	Humidity           float64
}

var weatherMap sync.Map
var regionMutex sync.Mutex

// simulateWeatherAPI simulates fetching weather data from a station.
func simulateWeatherAPI(stationID, region string) WeatherData {
	// Simulating data
	temperature := 20.0 + rand.Float64()*(30.0-20.0) // Random temperature between 20 and 30
	humidity := 30.0 + rand.Float64()*(70.0-30.0)    // Random humidity between 30 and 70
	return WeatherData{
		Temperature: temperature,
		Humidity:    humidity,
	}
}

// updateRegionData safely updates the weather data for a given region.
func updateRegionData(region string, data WeatherData) {
	regionMutex.Lock()
	defer regionMutex.Unlock()

	var regionData *RegionData
	if value, ok := weatherMap.Load(region); ok {
		regionData = value.(*RegionData)
	} else {
		// Region doesn't exist, initialize new data
		regionData = &RegionData{
			TemperatureBuffer: NewTemperatureRingBuffer(60), // Assuming 1 update per minute
		}
	}

	// Update temperature buffer and highest/lowest temperatures
	regionData.TemperatureBuffer.Add(data.Temperature)
	regionData.HighestTemperature = regionData.TemperatureBuffer.GetHighest()
	regionData.LowestTemperature = regionData.TemperatureBuffer.GetLowest()
	regionData.Humidity = data.Humidity

	weatherMap.Store(region, regionData)
}

// handleDynamicStations listens for new weather stations joining or leaving the system.
func handleDynamicStations(stationJoinCh <-chan string, stationLeaveCh <-chan string) {
	for {
		select {
		case stationID := <-stationJoinCh:
			region := stationID[:len(stationID)-len("-Station-1")] // Assuming region is the prefix of stationID
			fmt.Printf("Station %s joined, updating data for region %s\n", stationID, region)
			go func() {
				for {
					data := simulateWeatherAPI(stationID, region)
					updateRegionData(region, data)
					time.Sleep(time.Minute) // Simulate updates every minute
				}
			}()
		case stationID := <-stationLeaveCh:
			region := stationID[:len(stationID)-len("-Station-1")]
			fmt.Printf("Station %s left, updating data for region %s\n", stationID, region)
			// In a real implementation, you might need to handle station removal from the region's data.
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	stationJoinCh := make(chan string)
	stationLeaveCh := make(chan string)

	// Initialize some regions and stations for demonstration
	regions := []string{"Region A", "Region B"}
	for _, region := range regions {
		for i := 1; i <= 3; i++ {
			stationID := fmt.Sprintf("%s-Station-%d", region, i)
			stationJoinCh <- stationID
		}
	}

	// Start the goroutine to handle dynamic station joining and leaving
	go handleDynamicStations(stationJoinCh, stationLeaveCh)

	// Simulate new stations joining and leaving after a while
	time.Sleep(time.Second * 10)
	stationJoinCh <- "Region C-Station-1"
	time.Sleep(time.Second * 5)
	stationLeaveCh <- "Region B-Station-2"

	// Wait for some updates before printing the data
	time.Sleep(time.Minute * 5)

	fmt.Println("\nFinal weather data for each region:")
	weatherMap.Range(func(region, data interface{}) bool {
		regionData := data.(*RegionData)
		fmt.Printf("Region: %s\n", region)
		fmt.Printf("  Highest Temperature: %.2f°C\n", regionData.HighestTemperature)
		fmt.Printf("  Lowest Temperature: %.2f°C\n", regionData.LowestTemperature)
		fmt.Printf("  Current Humidity: %.2f%%\n", regionData.Humidity)
		return true
	})
}  