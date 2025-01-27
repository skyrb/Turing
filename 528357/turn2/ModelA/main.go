package main

import (
	"fmt"
	"strings"
)

// UserHealthData represents health data associated with a user
type UserHealthData struct {
	HeartRate    int
	ActivityLevel string
	Email        string // Include email field for deduplication
}

// Function to normalize an email address
func normalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

// Function to generate a unique ID with source prefix and email
func generateUserID(source, email string) string {
	return fmt.Sprintf("%s:%s", source, normalizeEmail(email))
}

// Function to merge data ensuring unique user IDs
func mergeData(userData map[string]UserHealthData, source string, newData map[string]UserHealthData) {
	for _, data := range newData {
		// Normalize the email and form a unique ID
		uniqueID := generateUserID(source, data.Email)

		if _, exists := userData[uniqueID]; exists {
			// Handle potential data merge if email exists across data sources
			// For now, we simply check and keep the newer data
			fmt.Printf("Duplicate found for ID %s. Overwriting data.\n", uniqueID)
			userData[uniqueID] = data
		} else {
			// No conflict, add new entry
			userData[uniqueID] = data
		}
	}
}

func main() {
	// Initialize a map to store consolidated user health data
	userData := make(map[string]UserHealthData)

	// Simulated data from wearable source
	wearableData := map[string]UserHealthData{
		"john@example.com":   {HeartRate: 72, ActivityLevel: "Active", Email: "john@example.com"},
		"jane@sample.com":    {HeartRate: 65, ActivityLevel: "Sedentary", Email: "jane@sample.com"},
		"john.doe@example.com": {HeartRate: 80, ActivityLevel: "Active", Email: "john.doe@example.com"},
	}

	// Simulated data from a health app with a duplication
	healthAppData := map[string]UserHealthData{
		"jane.doe@sample.com": {HeartRate: 67, ActivityLevel: "Moderately Active", Email: "jane.doe@sample.com"},
		"john@example.com":  {HeartRate: 75, ActivityLevel: "Active", Email: "john@example.com"}, // Possible duplicate
	}

	// Merge data from both sources ensuring unique user IDs
	mergeData(userData, "fitbit", wearableData)
	mergeData(userData, "healthapp", healthAppData)

	// Output the consolidated map
	fmt.Println("Consolidated user data:")
	for id, data := range userData {
		fmt.Printf("User ID: %s, Heart Rate: %d, Activity Level: %s, Email: %s\n", id, data.HeartRate, data.ActivityLevel, data.Email)
	}
}