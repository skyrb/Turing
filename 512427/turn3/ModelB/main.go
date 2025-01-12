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
	var rwMutex sync.RWMutex // Create a RWMutex to lock access to the map

	// Add some initial data for demonstration
	rwMutex.Lock()
	engagementMetrics["user123"] = &UserEngagement{Clicks: 2, TimeSpent: 5 * time.Minute}
	engagementMetrics["user456"] = &UserEngagement{Clicks: 4, TimeSpent: 3 * time.Minute}
	rwMutex.Unlock()

	// Create goroutines to simulate concurrent updates and reads
	for i := 0; i < 100; i++ {
		go updateUserEngagement(engagementMetrics, &rwMutex, "user123")
		go readUserEngagement(engagementMetrics, &rwMutex, "user123")
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to complete

	// Display the final engagement metrics
	fmt.Println("Final Engagement Metrics:")
	for userID, data := range engagementMetrics {
		fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %v\n", userID, data.Clicks, data.TimeSpent)
	}
}

func updateUserEngagement(metrics map[string]*UserEngagement, rwMutex *sync.RWMutex, userID string) {
	rwMutex.Lock() // Write lock
	defer rwMutex.Unlock()

	// Check if user exists, otherwise initialize
	if data, exists := metrics[userID]; exists {
		data.Clicks++
	} else {
		metrics[userID] = &UserEngagement{Clicks: 1, TimeSpent: 0}
	}
}

func readUserEngagement(metrics map[string]*UserEngagement, rwMutex *sync.RWMutex, userID string) {
	rwMutex.RLock() // Read lock
	defer rwMutex.RUnlock()

	// Read engagement data for a single user
	if data, exists := metrics[userID]; exists {
		fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %v\n", userID, data.Clicks, data.TimeSpent)
	} else {
		fmt.Printf("User ID: %s not found.\n", userID)
	}
}