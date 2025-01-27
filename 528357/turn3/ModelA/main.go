package main

import (
	"fmt"
	"time"
)

// UserHealthData represents health data associated with a user
type UserHealthData struct {
	HeartRate    int
	ActivityLevel string
	Timestamp    time.Time
}

// ConsolidatedData represents a map to store user health data with composite keys
var consolidatedData = make(map[string]UserHealthData)

// generateCompositeKey generates a unique key for each user entry based on user ID and timestamp
func generateCompositeKey(userID int, timestamp time.Time) string {
	return fmt.Sprintf("%d_%s", userID, timestamp.Format("20060102150405"))
}

// addHealthData adds a new entry of health data for a user
func addHealthData(userID int, heartRate int, activityLevel string, timestamp time.Time) {
	key := generateCompositeKey(userID, timestamp)
	consolidatedData[key] = UserHealthData{
		HeartRate:    heartRate,
		ActivityLevel: activityLevel,
		Timestamp:    timestamp,
	}
}

// cleanupStaleData removes entries older than the specified retention duration
func cleanupStaleData(retentionDuration time.Duration) {
	now := time.Now()
	for key, data := range consolidatedData {
		if now.Sub(data.Timestamp) > retentionDuration {
			delete(consolidatedData, key)
		}
	}
}

func main() {
	// Example data
	addHealthData(1, 72, "Active", time.Now().Add(-48*time.Hour))
	addHealthData(1, 75, "Moderately Active", time.Now().Add(-24*time.Hour))
	addHealthData(2, 65, "Sedentary", time.Now().Add(-72*time.Hour))
	addHealthData(3, 80, "Active", time.Now())

	// Display original data
	fmt.Println("Original data:")
	for key, data := range consolidatedData {
		fmt.Printf("Key: %s, Data: %+v\n", key, data)
	}

	// Cleanup data older than 48 hours
	cleanupStaleData(48 * time.Hour)

	// Display data after cleanup
	fmt.Println("\nData after cleanup:")
	for key, data := range consolidatedData {
		fmt.Printf("Key: %s, Data: %+v\n", key, data)
	}
}