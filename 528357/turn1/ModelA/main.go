package main

import (
	"fmt"
)

// UserHealthData represents health data associated with a user
type UserHealthData struct {
	HeartRate    int
	ActivityLevel string
}

// main function
func main() {
	// Initialize a map to store user health data
	userData := make(map[int]UserHealthData)

	// Simulated data from wearable source
	wearableData := map[int]UserHealthData{
		1: {HeartRate: 72, ActivityLevel: "Active"},
		2: {HeartRate: 65, ActivityLevel: "Sedentary"},
		3: {HeartRate: 80, ActivityLevel: "Active"},
	}

	// Simulated data from a health app
	healthAppData := map[int]UserHealthData{
		2: {HeartRate: 67, ActivityLevel: "Moderately Active"},
		4: {HeartRate: 75, ActivityLevel: "Active"},
	}

	// Function to merge data with unique user IDs
	mergeData := func(sourceData map[int]UserHealthData) {
		for id, data := range sourceData {
			originalID := id
			// Check if ID already exists in the userData map
			for {
				if _, exists := userData[id]; !exists {
					// If ID doesn't exist, add the data
					userData[id] = data
					break
				}
				// If ID exists, adjust ID until a unique ID is found
				id++
			}
			fmt.Printf("Imported user ID %d as %d\n", originalID, id)
		}
	}

	// Merge both data sources while ensuring unique keys
	mergeData(wearableData)
	mergeData(healthAppData)

	// Output the consolidated map
	fmt.Println("Consolidated user data:")
	for id, data := range userData {
		fmt.Printf("User ID: %d, Heart Rate: %d, Activity Level: %s\n", id, data.HeartRate, data.ActivityLevel)
	}
}