package main

import (
	"fmt"
)

// City represents a city with a name and population.
type City struct {
	Name       string
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

	city, err := findCityWithHighestPopulation(cities)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("City with the highest population: %s with %d people\n", city.Name, city.Population)
	}
}

// calculateTotalPopulation sums up the populations of the given cities.
func calculateTotalPopulation(cities []City) int {
	total := 0
	for _, city := range cities {
		total += city.Population
	}
	return total
}

// findCityWithHighestPopulation identifies the city with the highest population.
// It handles edge cases like an empty dataset and cities with the same population.
func findCityWithHighestPopulation(cities []City) (City, error) {
	if len(cities) == 0 {
		return City{}, fmt.Errorf("no cities available in the dataset")
	}

	highestPopulationCity := cities[0]
	for _, city := range cities {
		if city.Population > highestPopulationCity.Population {
			highestPopulationCity = city
		}
	}

	return highestPopulationCity, nil
}