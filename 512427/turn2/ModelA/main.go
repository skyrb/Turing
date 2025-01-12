package main

import (
	"fmt"
	"sync"
	"time"
)

type UserEngagement struct {
	Clicks int
	TimeSpent time.Duration
}

func main() {
	var engagementMetrics map[string]*UserEngagement = make(map[string]*UserEngagement)
	var mutex sync.Mutex // Create a mutex to lock access to the map

	// Add some initial data for demonstration
	engagementMetrics["user123"] = &UserEngagement{Clicks: 2, TimeSpent: 5 * time.Minute}
	engagementMetrics["user456"] = &UserEngagement{Clicks: 4, TimeSpent: 3 * time.Minute}

	// Create goroutines to simulate concurrent updates
	for i := 0; i < 100; i++ {
		go updateUserEngagement(engagementMetrics, &mutex, "user123", "user456")
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to complete

	// Display the final engagement metrics
	fmt.Println("Final Engagement Metrics:")
	for userID, data := range engagementMetrics {
		fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %v\n", userID, data.Clicks, data.TimeSpent)
	}
}

func updateUserEngagement(metrics map[string]*UserEngagement, mutex *sync.Mutex, userIDs ...string) {
	for _, userID := range userIDs {
		// Acquire the lock
		mutex.Lock()

		// Check if user exists, otherwise initialize
		if data, exists := metrics[userID]; exists {
			data.Clicks++
		} else {
			metrics[userID] = &UserEngagement{Clicks: 1, TimeSpent: 0}
		}

		// Release the lock
		mutex.Unlock()
	}
}