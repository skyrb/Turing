package main

import (
	"fmt"
	"sync"
)

// WaterQuality represents the live data from sensors
type WaterQuality struct {
	phLevel        float64
	turbidity      float64
	contamination  float64
}

// WaterTank represents a water tank with various monitoring parameters
type WaterTank struct {
	ID          string
	WaterLevel  float64
	SupplyArea  string
	QualityData WaterQuality
}

var (
	// waterTanks is a map from tank ID to WaterTank struct, protected by a mutex
	waterTanks   = make(map[string]WaterTank)
	tanksMutex   sync.RWMutex
)

// Function to update water level for a given tank ID
func updateWaterLevel(tankID string, newLevel float64) {
	tanksMutex.Lock()
	defer tanksMutex.Unlock()

	if tank, ok := waterTanks[tankID]; ok {
		tank.WaterLevel = newLevel
		waterTanks[tankID] = tank
		fmt.Println("Water level updated for Tank", tankID, "to", newLevel)
	} else {
		fmt.Println("Tank with ID", tankID, "not found.")
	}
}

// Function to update water quality for a given tank ID
func updateWaterQuality(tankID string, ph, turbidity, contamination float64) {
	tanksMutex.Lock()
	defer tanksMutex.Unlock()

	if tank, ok := waterTanks[tankID]; ok {
		tank.QualityData = WaterQuality{phLevel: ph, turbidity: turbidity, contamination: contamination}
		waterTanks[tankID] = tank
		fmt.Println("Water quality updated for Tank", tankID)
	} else {
		fmt.Println("Tank with ID", tankID, "not found.")
	}
}

// Function to retrieve water data for a given tank ID
func retrieveWaterData(tankID string) {
	tanksMutex.RLock()
	defer tanksMutex.RUnlock()

	if tank, ok := waterTanks[tankID]; ok {
		quality := tank.QualityData
		fmt.Printf("Water data for Tank %s - Level: %.1f, pH: %.2f, Turbidity: %.2f, Contamination: %.2f\n",
			tankID, tank.WaterLevel, quality.phLevel, quality.turbidity, quality.contamination)
	} else {
		fmt.Println("Tank with ID", tankID, "not found.")
	}
}

// Example usage:
func main() {
	// Initialize sample data
	tanksMutex.Lock()
	waterTanks["Tank1"] = WaterTank{ID: "Tank1", WaterLevel: 80.5, SupplyArea: "ZoneA"}
	waterTanks["Tank2"] = WaterTank{ID: "Tank2", WaterLevel: 75.2, SupplyArea: "ZoneB"}
	tanksMutex.Unlock()

	// Simulate concurrent updates
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		updateWaterLevel("Tank1", 70.0)
	}()

	go func() {
		defer wg.Done()
		updateWaterQuality("Tank2", 7.2, 0.5, 0.1)
	}()

	wg.Wait()

	retrieveWaterData("Tank1")
	retrieveWaterData("Tank2")
}