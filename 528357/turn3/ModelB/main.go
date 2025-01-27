package main

import (
	"fmt"
	"strings"
	"time"
)

// UserHealthData represents health data associated with a user at a specific timestamp
type UserHealthData struct {
	HeartRate    int
	ActivityLevel string
	Timestamp     time.Time
}

// main function
func main() {
	// Initialize a map to store user health data using a unique key (userID_timestamp)
	userData := make(map[string]UserHealthData)

	// Simulated data from wearable source
	wearableData := map[int]map[time.Time]UserHealthData{
		1: {
			time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC): {HeartRate: 72, ActivityLevel: "Active"},
			time.Date(2023, time.July, 16, 0, 0, 0, 0, time.UTC): {HeartRate: 78, ActivityLevel: "Active"},
		},
		2: {
			time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC): {HeartRate: 65, ActivityLevel: "Sedentary"},
		},
		3: {
			time.Date(2023, time.July, 16, 0, 0, 0, 0, time.UTC): {HeartRate: 80, ActivityLevel: "Active"},
		},
	}

	// Simulated data from a health app
	healthAppData := map[int]map[time.Time]UserHealthData{
		2: {
			time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC): {HeartRate: 67, ActivityLevel: "Moderately Active"},
		},
		4: {
			time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC): {HeartRate: 75, ActivityLevel: "Active"},
		},
	}

	// Function to generate a unique key for a user
	generateKey := func(userID int, timestamp time.Time) string {
		// Convert timestamp to a string in a unique format
		ts := timestamp.Format("2006-01-02T15:04:05")
		return fmt.Sprintf("%d_%s", userID, ts)
	}

	// Function to merge data while ensuring unique keys
	mergeData := func(sourceData map[int]map[time.Time]UserHealthData, sourcePrefix string) {
		for userID, dataByTimestamp := range sourceData {
			for timestamp, data := range dataByTimestamp {
				// Generate a unique key for the user
				key := generateKey(userID, timestamp)
				// Store the data in the consolidated map using the unique key
				userData[key] = UserHealthData{
					HeartRate:    data.HeartRate,
					ActivityLevel: data.ActivityLevel,
					Timestamp:    timestamp,
				}

				fmt.Printf("Imported user ID: %d, Timestamp: %s, Heart Rate: %d\n", userID, timestamp.Format("2006-01-02 15:04:05"), data.HeartRate)
			}
		}
	}

	// Merge both data sources while ensuring unique keys
	mergeData(wearableData, "fitbit")
	mergeData(healthAppData, "healthapp")

	// Function to clean up stale data (e.g., keep data only for the last 30 days)
	cleanupStaleData := func(data map[string]UserHealthData, retentionDays int) {
		currentTime := time.Now()