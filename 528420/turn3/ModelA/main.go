package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Simulates user activity data. Each entry contains a time spent value, in seconds.
type UserActivity struct {
	TimeSpent int64
}

// Metrics to track various statistics.
type Metrics struct {
	activeUsers     int32    // Current count of active users
	totalTimeSpent  int64    // Total time spent by users
	peakUsers       int32    // Peak number of active users
	mu              sync.Mutex // For safely accessing peakUsers
}

// processUserActivity simulates the processing of user activity data.
func processUserActivity(userChannel <-chan UserActivity, wg *sync.WaitGroup, metrics *Metrics) {
	defer wg.Done()
	for activity := range userChannel {
		// Increment active users and total time spent safely
		atomic.AddInt32(&metrics.activeUsers, 1)
		atomic.AddInt64(&metrics.totalTimeSpent, activity.TimeSpent)

		// Try to update peak user count if current active users exceed known peak
		// Use a mutex lock for safely updating peakUsers as it's beyond simple atomic operations.
		currentUsers := atomic.LoadInt32(&metrics.activeUsers)
		metrics.mu.Lock()
		if currentUsers > metrics.peakUsers {
			metrics.peakUsers = currentUsers
		}
		metrics.mu.Unlock()
	}

	// After processing, decrement active Users
	atomic.AddInt32(&metrics.activeUsers, -1)
}

func main() {
	metrics := &Metrics{
		activeUsers:    0,
		totalTimeSpent: 0,
		peakUsers:      0,
	}

	var wg sync.WaitGroup

	// Creating a channel to simulate user activity data stream
	userChannel := make(chan UserActivity, 100)

	// Number of goroutines to spawn
	numGoroutines := 5

	// Start multiple goroutines to process data from userChannel
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processUserActivity(userChannel, &wg, metrics)
	}

	// Simulate user activity with random time spent
	for i := 0; i < 100; i++ {
		timeSpent := int64(1 + i%10) // Simulating varying time spent
		userChannel <- UserActivity{TimeSpent: timeSpent}
	}

	// Close the channel to signal goroutines to stop processing
	close(userChannel)

	// Wait for all goroutines to finish
	wg.Wait()

	// Calculate average session time
	var averageSessionTime float64
	totalActiveSessions := 100
	if totalActiveSessions > 0 {
		averageSessionTime = float64(metrics.totalTimeSpent) / float64(totalActiveSessions)
	}

	// Output results
	fmt.Printf("Total Active Users: %d\n", metrics.activeUsers)
	fmt.Printf("Total Time Spent by Users: %d seconds\n", metrics.totalTimeSpent)
	fmt.Printf("Average Session Time: %.2f seconds\n", averageSessionTime)
	fmt.Printf("Peak Users: %d\n", metrics.peakUsers)
}