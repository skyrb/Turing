package main

import (
	"fmt"
)

// Function to count votes by region
func countVotesByRegion(data map[string]map[string]int) (map[string]map[string]int, map[string]int) {
	// Map to store total votes per option across all regions
	totalVotes := make(map[string]int)

	// Iterate through the data by region
	for region, optionsMap := range data {
		// Initialize a map for counts in the current region if needed
		if _, exists := data[region]; !exists {
			data[region] = make(map[string]int)
		}

		// Count votes for each option and aggregate to total votes
		for option, votes := range optionsMap {
			totalVotes[option] += votes
		}
	}
	return data, totalVotes
}

// Function to determine the most popular option by region
func mostPopularOptionByRegion(regionData map[string]map[string]int) map[string]string {
	popularOptions := make(map[string]string)

	for region, optionsMap := range regionData {
		maxVotes := -1
		popularOption := ""

		for option, votes := range optionsMap {
			if votes > maxVotes {
				maxVotes = votes
				popularOption = option
			}
		}
		popularOptions[region] = popularOption
	}

	return popularOptions
}

// Function to find the least popular option overall
func leastPopularOptionOverall(totalVotes map[string]int) string {
	minVotes := int(^uint(0) >> 1) // Set to max int
	leastPopular := ""

	for option, votes := range totalVotes {
		if votes < minVotes {
			minVotes = votes
			leastPopular = option
		}
	}

	return leastPopular
}

func main() {
	// Sample civic engagement data from multiple regions
	surveyData := map[string]map[string]int{
		"Region1": {"Option A": 200, "Option B": 150},
		"Region2": {"Option A": 300, "Option B": 250, "Option C": 100},
		"Region3": {"Option C": 400, "Option A": 100, "Option B": 50},
	}

	regionVotes, totalVotes := countVotesByRegion(surveyData)

	// Generate insights
	popularOptions := mostPopularOptionByRegion(regionVotes)
	leastPopularOption := leastPopularOptionOverall(totalVotes)

	// Print the most popular option by region
	fmt.Println("Most popular option by region:")
	for region, option := range popularOptions {
		fmt.Printf("%s: %s\n", region, option)
	}

	// Print the least popular option across all regions
	fmt.Printf("\nLeast popular option overall: %s\n", leastPopularOption)
}