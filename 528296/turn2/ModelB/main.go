package main  
import (  
    "fmt"
)

type City struct {
    Name       string
    Population int
}

func main() {
    // Test case 1: Normal case
    cities := []City{
        {Name: "New York", Population: 8804190},
        {Name: "Los Angeles", Population: 3898747},
        {Name: "Chicago", Population: 2746388},
    }
    highestPopCity, totalPopulation := findHighestPopCity(cities)
    fmt.Println("Highest population city:", highestPopCity.Name)
    fmt.Println("Total population across cities:", totalPopulation)

    // Test case 2: Empty dataset
    emptyCities := []City{}
    highestPopCity, totalPopulation = findHighestPopCity(emptyCities)
    fmt.Println("\nHighest population city (empty dataset):", highestPopCity.Name)
    fmt.Println("Total population across cities:", totalPopulation)

    // Test case 3: Cities with the same population
    citiesWithSamePop := []City{
        {Name: "City A", Population: 10000},
        {Name: "City B", Population: 10000},
    }
    highestPopCity, totalPopulation = findHighestPopCity(citiesWithSamePop)
    fmt.Println("\nHighest population city (same population):", highestPopCity.Name)
    fmt.Println("Total population across cities:", totalPopulation)
}

func findHighestPopCity(cities []City) (City, int) {
    if len(cities) == 0 {
        // Handle empty dataset edge case
        return City{}, 0
    }

    highestPopCity := cities[0]
    totalPopulation := highestPopCity.Population

    for _, city := range cities[1:] {
        if city.Population > highestPopCity.Population {
            highestPopCity = city
        } else if city.Population == highestPopCity.Population {
            // Handle multiple cities with the same population
            // Here, we pick the city first alphabetically
            highestPopCity = City{Name: min(highestPopCity.Name, city.Name), Population: city.Population}
        }
        totalPopulation += city.Population
    }

    return highestPopCity, totalPopulation
}

// Function to find the minimum of two strings
func min(a, b string) string {
    if a < b {
        return a
    }
    return b
}
