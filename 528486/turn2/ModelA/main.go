package main

import (
	"fmt"
	"time"
)

// Function to parse timestamps and count votes with timestamps
func analyzeEngagement(surveyData map[string]map[string]string) map[string]map[string]int {
	// Create a map to store timestamps count for each option
	engagementTrends := make(map[string]map[string]int)

	// Iterate through the survey data
	for _, options := range surveyData {
		// Iterate through each user's options with timestamps
		for option, timestamp := range options {
			// Initialize the map for the option if not already present
			if _, exists := engagementTrends[option]; !exists {
				engagementTrends[option] = make(map[string]int)
			}
			
			// Parse the timestamp to a more readable hour-minute format
			parsedTime, err := time.Parse("2006-01-02 15:04", timestamp)
			if err != nil {
				fmt.Println("Error parsing time:", err)
				continue
			}
			formattedTime := parsedTime.Format("15:04")
			
			// Increment the count for the timestamp
			engagementTrends[option][formattedTime]++
		}
	}

	return engagementTrends
}

// Function to find the peak time for each option
func findPeakTimes(engagementTrends map[string]map[string]int) map[string]string {
	peakTimes := make(map[string]string)

	// Iterate over each option and its timestamp counts
	for option, times := range engagementTrends {
		maxCount := 0
		peakTime := ""

		// Determine the timestamp with the highest count
		for time, count := range times {
			if count > maxCount {
				maxCount = count
				peakTime = time
			}
		}

		peakTimes[option] = peakTime
	}

	return peakTimes
}

func main() {
	// Sample survey data with timestamps
	surveyData := map[string]map[string]string{
		"User1": {"Option A": "2025-01-01 10:00", "Option B": "2025-01-01 10:30"},
		"User2": {"Option B": "2025-01-01 11:00", "Option C": "2025-01-01 11:15"},
		"User3": {"Option A": "2025-01-01 10:00", "Option B": "2025-01-01 11:00"},
		"User4": {"Option A": "2025-01-01 10:30"},
	}

	// Analyze engagement trends
	engagementTrends := analyzeEngagement(surveyData)

	// Find peak engagement times
	peakTimes := findPeakTimes(engagementTrends)

	// Print the peak engagement times
	fmt.Println("Peak engagement times for each option:")
	for option, time := range peakTimes {
		fmt.Printf("%s: %s\n", option, time)
	}
}