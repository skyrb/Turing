package main

import (
	"fmt"
	"strings"
)

// UserHealthData represents health data associated with a user
type UserHealthData struct {
	HeartRate    int
	ActivityLevel string
}

// main function
func main() {
	// Initialize a map to store user health data using a unique key
	userData := make(map[string]UserHealthData)

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

	// Function to generate a unique key for a user
	generateKey := func(email string, sourcePrefix string) string {
		// Convert email to lowercase for case-insensitive comparison
		email = strings.ToLower(email)
		// Return the combination of source prefix and email
		return fmt.Sprintf("%s:%s", sourcePrefix, email)
	}

	// Function to merge data while ensuring unique keys
	mergeData := func(sourceData map[int]UserHealthData, sourcePrefix string) {
		for id, data := range sourceData {
			// Get the email address associated with the user ID
			email := fmt.Sprintf("user%d@example.com", id)

			// Generate a unique key for the user
			key := generateKey(email, sourcePrefix)

			// Store the data in the consolidated map using the unique key
			userData[key] = data

			fmt.Printf("Imported user %s as %s\n", email, key)
		}
	}

	// Merge both data sources while ensuring unique keys
	mergeData(wearableData, "fitbit")
	mergeData(healthAppData, "healthapp")

	// Output the consolidated map
	fmt.Println("\nConsolidated user data:")
	for key, data := range userData {
		fmt.Printf("User: %s, Heart Rate: %d, Activity Level: %s\n", key, data.HeartRate, data.ActivityLevel)
	}
}
 