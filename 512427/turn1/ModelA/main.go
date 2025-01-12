package main

import (
	"fmt"
)

// UserEngagement struct to hold user engagement metrics
type UserEngagement struct {
	Clicks int    `json:"clicks"`
	Time   int64  `json:"time"` // Time spent in seconds
}

func main() {
	// Create a map to store user engagement data
	userEngagementMap := make(map[string]UserEngagement)

	// Function to add or update user engagement data
	addOrUpdateEngagement := func(userID string, clicks int, timeSpent int64) {
		// Check if the user already exists in the map
		if existingEngagement, ok := userEngagementMap[userID]; ok {
			// Update the existing user's metrics
			existingEngagement.Clicks += clicks
			existingEngagement.Time += timeSpent
			userEngagementMap[userID] = existingEngagement
		} else {
			// Add the new user to the map
			newEngagement := UserEngagement{Clicks: clicks, Time: timeSpent}
			userEngagementMap[userID] = newEngagement
		}
	}

	// Example usage: add or update engagement data for users
	addOrUpdateEngagement("user1", 2, 120) // User1 clicked twice and spent 120 seconds on a page
	addOrUpdateEngagement("user2", 5, 300) // User2 clicked five times and spent 300 seconds on a page
	addOrUpdateEngagement("user1", 1, 180) // Update User1's data: clicked once more and spent an additional 180 seconds

	// Print the user engagement data
	for userID, engagement := range userEngagementMap {
		fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %ds\n", userID, engagement.Clicks, engagement.Time)
	}
}