package main  
import "fmt"

type WaterTank struct {
    ID         string
    WaterLevel float64
    SupplyArea string
}

func main() {
    // Initialize the waterTanks map as an empty map from tank ID to WaterTank struct
    waterTanks := make(map[string]WaterTank)

    // Add sample data for water tanks
    waterTanks["Tank1"] = WaterTank{ID: "Tank1", WaterLevel: 80.5, SupplyArea: "ZoneA"}
    waterTanks["Tank2"] = WaterTank{ID: "Tank2", WaterLevel: 75.2, SupplyArea: "ZoneB"}
    waterTanks["Tank3"] = WaterTank{ID: "Tank3", WaterLevel: 90.3, SupplyArea: "ZoneA"}

    // Function to update water level for a given tank ID
    updateWaterLevel := func(tankID string, newLevel float64) {
        if tank, ok := waterTanks[tankID]; ok {
            tank.WaterLevel = newLevel
            waterTanks[tankID] = tank
            fmt.Println("Water level updated for Tank", tankID, "to", newLevel)
        } else {
            fmt.Println("Tank with ID", tankID, "not found.")
        }
    }

    // Function to retrieve water level for a given tank ID
    retrieveWaterLevel := func(tankID string) {
        if tank, ok := waterTanks[tankID]; ok {
            fmt.Println("Water level for Tank", tankID, "is:", tank.WaterLevel)
        } else {
            fmt.Println("Tank with ID", tankID, "not found.")
        }
    }
    
    // Function to find tanks with water level below a given threshold
    findShortageTanks := func(threshold float64) {
        for _, tank := range waterTanks {
            if tank.WaterLevel < threshold {
                fmt.Println("Tank", tank.ID, "in Supply Area", tank.SupplyArea, "is short of water.")
            }
        }
    }
    
    // Example usage:
    updateWaterLevel("Tank2", 68.0)
    retrieveWaterLevel("Tank3")
    findShortageTanks(80.0)
}