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

// HealthDataStore represents a nested map structure to store user health data
var healthDataStore = make(map[int]map[string]UserHealthData)

// addHealthData adds a new entry of health data for a user
func addHealthData(userID int, heartRate int, activityLevel string, timestamp time.Time) {
	// Format timestamp as a string key
	timeKey := timestamp.Format("20060102150405")

	// Initialize the inner map if it doesn't exist
	if _, exists := healthDataStore[userID]; !exists {
		healthDataStore[userID] = make(map[string]UserHealthData)
	}

	// Add the health data entry
	healthDataStore[userID][timeKey] = UserHealthData{
		HeartRate:    heartRate,
		ActivityLevel: activityLevel,
		Timestamp:    timestamp,
	}
}

// getUserData returns all historical entries for a specific user
func getUserData(userID int) []UserHealthData {
	var userData []UserHealthData
	// Check if the user exists in the store
	if entries, exists := healthDataStore[userID]; exists {
		for _, data := range entries {
			userData = append(userData, data)
		}
	}
	return userData
}

func main() {
	// Example data
	addHealthData(1, 72, "Active", time.Now().Add(-48*time.Hour))
	addHealthData(1, 75, "Moderately Active", time.Now().Add(-24*time.Hour))
	addHealthData(2, 65, "Sedentary", time.Now().Add(-72*time.Hour))
	addHealthData(3, 80, "Active", time.Now())

	// Query for all historical entries for a specific user
	userID := 1
	fmt.Printf("Historical data for user %d:\n", userID)
	userData := getUserData(userID)
	for _, data := range userData {
		fmt.Printf("Timestamp: %s, Heart Rate: %d, Activity Level: %s\n",
			data.Timestamp.Format("2006-01-02 15:04:05"), data.HeartRate, data.ActivityLevel)
	}
}