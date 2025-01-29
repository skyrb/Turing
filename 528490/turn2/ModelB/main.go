package main

import (
	"fmt"
	"sync"
	"time"
)

type WaterTank struct {
	ID         string
	WaterLevel float64
	SupplyArea string
	SensorData *SensorData
}

type SensorData struct {
	pH         float64
	Turbidity  float64
	Contamination float64
}

func main() {
	// Initialize the waterTanks map
	waterTanks := make(map[string]WaterTank)

	// Add sample data for water tanks
	waterTanks["Tank1"] = WaterTank{ID: "Tank1", WaterLevel: 80.5, SupplyArea: "ZoneA", SensorData: &SensorData{pH: 7.2, Turbidity: 12.5, Contamination: 0.3}}
	waterTanks["Tank2"] = WaterTank{ID: "Tank2", WaterLevel: 75.2, SupplyArea: "ZoneB", SensorData: &SensorData{pH: 6.8, Turbidity: 8.9, Contamination: 0.2}}
	waterTanks["Tank3"] = WaterTank{ID: "Tank3", WaterLevel: 90.3, SupplyArea: "ZoneA", SensorData: &SensorData{pH: 7.5, Turbidity: 10.0, Contamination: 0.4}}

	// Create a channel to receive sensor data updates
	sensorDataUpdates := make(chan SensorUpdate)

	// Go routine to handle sensor data updates
	go func() {
		for update := range sensorDataUpdates {
			if tank, ok := waterTanks[update.TankID]; ok {
				// Update sensor data atomically using a mutex
				var mu sync.Mutex
				mu.Lock()
				tank.SensorData = update.SensorData
				waterTanks[update.TankID] = tank
				mu.Unlock()
				fmt.Println("Sensor data updated for Tank", update.TankID)
			} else {
				fmt.Println("Tank with ID", update.TankID, "not found for sensor data update.")
			}
		}
	}()

	// Simulate sensor data updates
	for {
		go sendSensorUpdate(sensorDataUpdates, "Tank1", &SensorData{pH: 7.3, Turbidity: 13.0, Contamination: 0.35})
		go sendSensorUpdate(sensorDataUpdates, "Tank2", &SensorData{pH: 6.75, Turbidity: 9.2, Contamination: 0.18})
		go sendSensorUpdate(sensorDataUpdates, "Tank4", &SensorData{pH: 7.4, Turbidity: 11.5, Contamination: 0.42}) // Tank4 doesn't exist, to test error handling
		time.Sleep(2 * time.Second)
	}
}

// Function to send sensor data updates asynchronously using a goroutine
func sendSensorUpdate(updates chan<- SensorUpdate, tankID string, sensorData *SensorData) {
	updates <- SensorUpdate{TankID: tankID, SensorData: sensorData}
}

// Structure to hold sensor data update information
type SensorUpdate struct {
	TankID     string
	SensorData *SensorData
}
 