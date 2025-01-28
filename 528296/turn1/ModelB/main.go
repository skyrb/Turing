
package main

import (
    "fmt"
)
// define a City struct
type City struct {
    Name     string
    Population int
}

func main() {
    // Example civic tech dataset: a list of cities with their populations
    cities := []City{
        {Name: "New York", Population: 8804190},
        {Name: "Los Angeles", Population: 3898747},
        {Name: "Chicago", Population: 2746388},
        {Name: "Houston", Population: 2304580},
        {Name: "Phoenix", Population: 1608139},
    }

    totalPopulation := calculateTotalPopulation(cities)
    fmt.Println("Total population across cities:", totalPopulation)
}

// CalculateTotalPopulation: This function uses a range loop to iterate 
// through the slice of cities and sum up their populations.
func calculateTotalPopulation(cities []City) int {
    total := 0
    for _, city := range cities {
        total += city.Population
    }
    return total
}
