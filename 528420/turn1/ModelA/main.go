package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// processUserActivity simulates processing activity data from users.
// It takes a channel of user activity data and a pointer to the active user counter.
func processUserActivity(userChannel <-chan struct{}, wg *sync.WaitGroup, activeUsers *int32) {
	defer wg.Done()
	for range userChannel {
		// Increment the shared active users counter atomically
		atomic.AddInt32(activeUsers, 1)
	}
}

func main() {
	var activeUsers int32
	var wg sync.WaitGroup

	// Create a channel to simulate user activity data stream
	userChannel := make(chan struct{}, 100)

	// Number of goroutines to spawn
	numGoroutines := 5

	// Start multiple goroutines to process data from userChannel
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processUserActivity(userChannel, &wg, &activeUsers)
	}

	// Simulate user activity
	for i := 0; i < 100; i++ {
		userChannel <- struct{}{}
	}

	// Close the channel to signal goroutines to stop processing
	close(userChannel)

	// Wait for all goroutines to finish
	wg.Wait()

	// Output the total count of active users
	fmt.Printf("Total Active Users: %d\n", activeUsers)
}