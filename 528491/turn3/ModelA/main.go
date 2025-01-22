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
	Timestamp   time.Time
}

// RegionAnalytics holds the analytics data for each region
type RegionAnalytics struct {
	Temperatures []float64 // Maintaining a list of temperatures
	Mutex        sync.Mutex // Ensuring thread-safe updates
}

// Weather System tracks data and analytics per region
type WeatherSystem struct {
	Data      sync.Map             // Stores latest weather data per region
	Analytics sync.Map             // Stores real-time analytics per region
	Updates   chan WeatherDataUpdate // Channel for incoming data updates
	quit      chan bool
}

// WeatherDataUpdate holds updated data from a station
type WeatherDataUpdate struct {
	Region string
	Data   WeatherData
}

// NewWeatherSystem initializes a new weather system
func NewWeatherSystem() *WeatherSystem {
	return &WeatherSystem{
		Updates: make(chan WeatherDataUpdate, 100),
		quit:    make(chan bool),
	}
}

// StartProcessing starts the processing of incoming weather data
func (ws *WeatherSystem) StartProcessing() {
	go func() {
		for {
			select {
			case update := <-ws.Updates:
				ws.handleUpdate(update)
			case <-ws.quit:
				return
			}
		}
	}()
}

// StopProcessing stops the data processing loop
func (ws *WeatherSystem) StopProcessing() {
	ws.quit <- true
}

// handleUpdate processes incoming weather data updates
func (ws *WeatherSystem) handleUpdate(update WeatherDataUpdate) {
	// Insert the latest data
	ws.Data.Store(update.Region, update.Data)

	// Update analytics
	value, _ := ws.Analytics.LoadOrStore(update.Region, &RegionAnalytics{
		Temperatures: []float64{}})
	analytics := value.(*RegionAnalytics)

	analytics.Mutex.Lock()
	analytics.Temperatures = append(analytics.Temperatures, update.Data.Temperature)
	// Remove data older than an hour
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	tempSlice := make([]float64, 0)
	for _, v := range analytics.Temperatures {
		if update.Data.Timestamp.After(oneHourAgo) {
			tempSlice = append(tempSlice, v)
		}
	}
	analytics.Temperatures = tempSlice
	analytics.Mutex.Unlock()
}

// getRealTimeStats retrieves processed statistics for a region
func (ws *WeatherSystem) getRealTimeStats(region string) (maxTemp, minTemp float64) {
	value, ok := ws.Analytics.Load(region)
	if !ok {
		return 0, 0 // Handle zero as default when region has no analytics
	}
	analytics := value.(*RegionAnalytics)

	analytics.Mutex.Lock()
	defer analytics.Mutex.Unlock()
	if len(analytics.Temperatures) == 0 {
		return 0, 0
	}

	maxTemp = analytics.Temperatures[0]
	minTemp = analytics.Temperatures[0]
	for _, temp := range analytics.Temperatures {
		if temp > maxTemp {
			maxTemp = temp
		}
		if temp < minTemp {
			minTemp = temp
		}
	}
	return
}

// simulateWeatherAPI simulates fetching weather data from an API.
func simulateWeatherAPI(stationID, region string) WeatherData {
	temperature := 20.0 + rand.Float64()*(30.0-20.0)
	humidity := 30.0 + rand.Float64()*(70.0-30.0)
	return WeatherData{
		Temperature: temperature,
		Humidity:    humidity,
		Timestamp:   time.Now(),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ws := NewWeatherSystem()
	ws.StartProcessing()
	defer ws.StopProcessing()

	regions := []string{"Region A", "Region B"}

	// Launch dynamic weather station simulations
	var wg sync.WaitGroup
	for _, region := range regions {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()
			stationID := 0
			for i := 0; i < 10; i++ {
				stationID++
				data := simulateWeatherAPI(fmt.Sprintf("%s-Station-%d", region, stationID), region)
				ws.Updates <- WeatherDataUpdate{Region: region, Data: data}
				time.Sleep(time.Second * 1) // Simulating periodic updates every second
			}
		}(region)
	}

	wg.Wait()

	// Get and print analytics for each region
	for _, region := range regions {
		maxTemp, minTemp := ws.getRealTimeStats(region)
		fmt.Printf("Real-time stats for %s: Max Temp: %.2f°C, Min Temp: %.2f°C\n", region, maxTemp, minTemp)
	}
}