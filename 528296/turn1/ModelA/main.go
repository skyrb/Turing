package main

import (
    "fmt"
)

// calculateTotalPopulation takes a map of city populations and returns the total population.
func calculateTotalPopulation(cityPopulations map[string]int) int {
    totalPopulation := 0

    // Range over the map to calculate the total population
    for _, population := range cityPopulations {
        totalPopulation += population
    }

    return totalPopulation
}

func main() {
    // Sample dataset of city populations
    cityPopulations := map[string]int{
        "New York":    8419000,
        "Los Angeles": 3980000,
        "Chicago":     2716000,
        "Houston":     2313000,
        "Phoenix":     1690000,
    }

    // Calculate the total population
    totalPopulation := calculateTotalPopulation(cityPopulations)

    // Print the total population
    fmt.Printf("The total population across cities is %d\n", totalPopulation)
}