package main  
import (  
    "fmt"
    "sync"
)

type City struct {
    Name       string
    Population int
}

type Dataset struct {
    Name     string
    Data     []City
}

func main() {
    // Create datasets
    populationDataset := Dataset{
        Name: "Population",
        Data: []City{
            {Name: "New York", Population: 8804190},
            {Name: "Los Angeles", Population: 3898747},
            {Name: "Chicago", Population: 2746388},
        },
    }

    infrastructureDataset := Dataset{
        Name: "Infrastructure",
        Data: []City{
            {Name: "New York", Population: 8804190},
            {Name: "Los Angeles", Population: 3898747},
            {Name: "Chicago", Population: 2746388},
        },
    }

    // Create a channel to receive results
    resultChannel := make(chan string)

    // Process multiple datasets concurrently
    var wg sync.WaitGroup
    wg.Add(2)

    go processDataset(populationDataset, resultChannel, &wg)
    go processDataset(infrastructureDataset, resultChannel, &wg)

    wg.Wait()

    close(resultChannel)

    // Print results
    for result := range resultChannel {
        fmt.Println(result)
    }
}

func processDataset(dataset Dataset, resultChannel chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    highestPopCity, totalPopulation := findHighestPopCity(dataset.Data)
    result := fmt.Sprintf("Highest population city in %s dataset: %s. Total population: %d", dataset.Name, highestPopCity.Name, totalPopulation)
    resultChannel <- result
}

func findHighestPopCity(cities []City) (City, int) {  
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