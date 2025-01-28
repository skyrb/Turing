package main

import (
	"fmt"
	"sync"
)

// City struct to hold population and infrastructure data
type City struct {
	Name            string
	Population      int
	Infrastructure  string
}

// PopulationData struct to represent population data from a dataset
type PopulationData struct {
	Name       string
	Population int
}

// InfrastructureData struct to represent infrastructure data from another dataset
type InfrastructureData struct {
	Name          string
	Infrastructure string
}

func main() {
	populationData := []PopulationData{
		{Name: "New York", Population: 8804190},
		{Name: "Los Angeles", Population: 3898747},
		{Name: "Chicago", Population: 2746388},
	}

	infrastructureData := []InfrastructureData{
		{Name: "New York", Infrastructure: "Advanced"},
		{Name: "Los Angeles", Infrastructure: "Intermediate"},
		{Name: "Chicago", Infrastructure: "Basic"},
	}

	var wg sync.WaitGroup

	cities := make(chan []City, 2)

	wg.Add(2)
	go processPopulationData(populationData, cities, &wg)
	go processInfrastructureData(infrastructureData, cities, &wg)

	wg.Wait()
	close(cities)

	mergedCities := mergeDatasets(cities)
	for _, city := range mergedCities {
		fmt.Printf("City: %s, Population: %d, Infrastructure: %s\n", city.Name, city.Population, city.Infrastructure)
	}
}

func processPopulationData(data []PopulationData, cities chan<- []City, wg *sync.WaitGroup) {
	defer wg.Done()

	cityList := make([]City, 0, len(data))
	for _, item := range data {
		cityList = append(cityList, City{Name: item.Name, Population: item.Population})
	}
	cities <- cityList
}

func processInfrastructureData(data []InfrastructureData, cities chan<- []City, wg *sync.WaitGroup) {
	defer wg.Done()

	cityList := make([]City, 0, len(data))
	for _, item := range data {
		cityList = append(cityList, City{Name: item.Name, Infrastructure: item.Infrastructure})
	}
	cities <- cityList
}

func mergeDatasets(cities <-chan []City) []City {
	cityMap := make(map[string]*City)

	for cityList := range cities {
		for _, city := range cityList {
			if existing, found := cityMap[city.Name]; found {
				if city.Population != 0 {
					existing.Population = city.Population
				}
				if city.Infrastructure != "" {
					existing.Infrastructure = city.Infrastructure
				}
			} else {
				cityCopy := city
				cityMap[city.Name] = &cityCopy
			}
		}
	}

	mergedCities := make([]City, 0, len(cityMap))
	for _, city := range cityMap {
		mergedCities = append(mergedCities, *city)
	}

	return mergedCities
}