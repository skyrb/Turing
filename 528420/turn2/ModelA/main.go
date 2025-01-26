package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Simulates user activity data. Each entry consists of a time spent value, in seconds.
type UserActivity struct {
	TimeSpent int64
}

// processUserActivity simulates processing activity data from users.
// It takes a channel of user activity data, a pointer to the active user counter,
// and a pointer to the total time spent counter.
func processUserActivity(userChannel <-chan UserActivity, wg *sync.WaitGroup, activeUsers *int32, totalTimeSpent *int64) {
	defer wg.Done()
	for activity := range userChannel {
		// Increment the shared active users counter atomically
		atomic.AddInt32(activeUsers, 1)

		// Add the time spent value to the total time spent atomically
		atomic.AddInt64(totalTimeSpent, activity.TimeSpent)
	}
}

func main() {
	var activeUsers int32
	var totalTimeSpent int64
	var wg sync.WaitGroup

	// Create a channel to simulate user activity data stream
	userChannel := make(chan UserActivity, 100)

	// Number of goroutines to spawn
	numGoroutines := 5

	// Start multiple goroutines to process data from userChannel
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processUserActivity(userChannel, &wg, &activeUsers, &totalTimeSpent)
	}

	// Simulate user activity with random time spent
	for i := 0; i < 100; i++ {
		// Simulate a random time spent between 1 and 10 seconds
		timeSpent := int64(1 + time.Now().UnixNano()%10)
		userChannel <- UserActivity{TimeSpent: timeSpent}
	}

	// Close the channel to signal goroutines to stop processing
	close(userChannel)

	// Wait for all goroutines to finish
	wg.Wait()

	// Output the results
	fmt.Printf("Total Active Users: %d\n", activeUsers)
	fmt.Printf("Total Time Spent by Users: %d seconds\n", totalTimeSpent)
}