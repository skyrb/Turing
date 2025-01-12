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

var engagementMetrics = make(map[string]*UserEngagement)
var engagementMutex sync.Mutex

func updateEngagement(userID string, clicks int) {
    engagementMutex.Lock()
    defer engagementMutex.Unlock()

    if data, exists := engagementMetrics[userID]; exists {
        data.Clicks += clicks
    } else {
        engagementMetrics[userID] = &UserEngagement{Clicks: clicks}
    }
}

func main() {
    // Initialize user engagement data
    engagementMetrics["user123"] = &UserEngagement{Clicks: 0}
    engagementMetrics["user456"] = &UserEngagement{Clicks: 0}

    // Start multiple goroutines to simulate concurrent updates
    numGoroutines := 10
    clicksPerGoroutine := 100

    for i := 0; i < numGoroutines; i++ {
        go func() {
            for _, userID := range []string{"user123", "user456"} {
                updateEngagement(userID, clicksPerGoroutine)
            }
        }()
    }

    // Wait for all goroutines to complete
    time.Sleep(time.Second * 2)

    // Display the final engagement metrics
    engagementMutex.Lock()
    defer engagementMutex.Unlock()
    for userID, data := range engagementMetrics {
        fmt.Printf("User ID: %s, Clicks: %d\n", userID, data.Clicks)
    }
} 
