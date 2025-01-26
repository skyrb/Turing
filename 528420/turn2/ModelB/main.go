package main  
import (  
    "fmt"
    "sync"
    "sync/atomic"
)

// UserActivity represents a single user activity event.
type UserActivity struct {
    IsActive bool
    TimeSpent int64 // In milliseconds
}

// processUserActivity handles user activity data and updates the shared counters.
func processUserActivity(userChannel <-chan UserActivity, wg *sync.WaitGroup, activeUsers *int32, totalTimeSpent *int64) {
    defer wg.Done()

    for activity := range userChannel {
        // Increment or decrement the active user counter based on activity.IsActive
        if activity.IsActive {
            atomic.AddInt32(activeUsers, 1)
        } else {
            atomic.AddInt32(activeUsers, -1)
        }

        // Add the time spent to the total
        atomic.AddInt64(totalTimeSpent, activity.TimeSpent)
    }
}

func main() {  
    var activeUsers int32
    var totalTimeSpent int64
    var wg sync.WaitGroup

    // User activity stream channel
    userChannel := make(chan UserActivity, 100)

    numGoroutines := 5

    // Start goroutines to process activities
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go processUserActivity(userChannel, &wg, &activeUsers, &totalTimeSpent)
    }

    // Simulate user activity, each activity has an indicator and a time spent
    activities := []UserActivity{
        {IsActive: true, TimeSpent: 1000}, // 1 second active
        {IsActive: true, TimeSpent: 500}, // 0.5 seconds active
        {IsActive: false, TimeSpent: 0}, // Inactive
        {IsActive: true, TimeSpent: 2000}, // 2 seconds active
        // Add more activities as needed
    }

    for _, activity := range activities {
        userChannel <- activity
    }

    // Close the channel to signal goroutines to finish
    close(userChannel)

    // Wait for all goroutines to finish
    wg.Wait()

    // Convert totalTimeSpent from milliseconds to seconds
    totalTimeSpentInSec := float64(totalTimeSpent) / 1000

    fmt.Printf("Total Active Users: %d\n", activeUsers)
    fmt.Printf("Total Time Spent by Active Users: %.2fs\n", totalTimeSpentInSec)
}