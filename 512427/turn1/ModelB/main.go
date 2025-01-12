
package main  
import (  
    "fmt"
    "time"
)

type UserEngagement struct {
    Clicks int
    TimeSpent time.Duration
}

func main() {  
    engagementMetrics := make(map[string]*UserEngagement)

    // Adding user engagement data
    engagementMetrics["user123"] = &UserEngagement{Clicks: 2, TimeSpent: 5 * time.Minute}
    engagementMetrics["user456"] = &UserEngagement{Clicks: 4, TimeSpent: 3 * time.Minute}

    // Updating user engagement data
    userID := "user123"
    if data, exists := engagementMetrics[userID]; exists {
        data.Clicks++
        data.TimeSpent += 2 * time.Minute
    } else {
        engagementMetrics[userID] = &UserEngagement{Clicks: 1, TimeSpent: 2 * time.Minute}
    }

    // Displaying the engagement metrics
    for userID, data := range engagementMetrics {
        fmt.Printf("User ID: %s, Clicks: %d, Time Spent: %v\n", userID, data.Clicks, data.TimeSpent)
    }
}
  