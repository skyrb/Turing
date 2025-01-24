package main

import (
	"fmt"
)

// Function to count the votes for each option
func countVotes(surveyData map[string][]string) map[string]int {
	// Create a map to store the count of each option
	optionCounts := make(map[string]int)

	// Iterate through the survey data
	for _, options := range surveyData {
		// Iterate through each user's selected options
		for _, option := range options {
			// Increment the count for the option
			optionCounts[option]++
		}
	}

	return optionCounts
}

func main() {
	// Sample survey data
	surveyData := map[string][]string{
		"User1": {"Option A", "Option B"},
		"User2": {"Option B", "Option C"},
		"User3": {"Option A", "Option C", "Option B"},
		"User4": {"Option A"},
	}

	// Count the votes for each option
	voteCounts := countVotes(surveyData)

	// Print the results
	fmt.Println("Vote counts for each option:")
	for option, count := range voteCounts {
		fmt.Printf("%s: %d\n", option, count)
	}
}