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
	var rwMutex sync.RWMutex // Create a read-write mutex to manage access

	// Add some initial data for demonstration
	engagementMetrics["user123"] = &UserEngagement{Clicks: 2, TimeSpent: 5 * time.Minute}
	engagementMetrics["user456"] = &UserEngagement{Clicks: 4, TimeSpent: 3 * time.Minute}

	// Create goroutines to simulate concurrent updates and reads
	for i := 0; i < 100; i++ {
		go updateUserEngagement(engagementMetrics, &rwMutex, "user123", "user456")
	}
	for i := 0; i < 1000; i++ {
		go readUserEngagement(engagementMetrics, &rwMutex)
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to complete

	// Display the final engagement metrics
	rwMutex.RLock()
	fmt.Println("Final Engagement Metrics:")
	for userID, data := range engagementMetrics {
		fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %v\n", userID, data.Clicks, data.TimeSpent)
	}
	rwMutex.RUnlock()
}

func updateUserEngagement(metrics map[string]*UserEngagement, rwMutex *sync.RWMutex, userIDs ...string) {
	for _, userID := range userIDs {
		// Acquire the write lock
		rwMutex.Lock()

		// Check if user exists, otherwise initialize
		if data, exists := metrics[userID]; exists {
			data.Clicks++
		} else {
			metrics[userID] = &UserEngagement{Clicks: 1, TimeSpent: 0}
		}

		// Release the write lock
		rwMutex.Unlock()
	}
}

func readUserEngagement(metrics map[string]*UserEngagement, rwMutex *sync.RWMutex) {
	// Acquire the read lock
	rwMutex.RLock()

	fmt.Println("Reading engagement metrics...")
	for userID, data := range metrics {
		fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %v\n", userID, data.Clicks, data.TimeSpent)
	}

	// Release the read lock
	rwMutex.RUnlock()
}