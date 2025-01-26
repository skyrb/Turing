package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// UserActivity represents user activity data.
type UserActivity struct {
	TimeSpent int64
	StartTime time.Time
}

// Session represents a user's session.
type Session struct {
	TotalTimeSpent int64
	StartTime      time.Time
}

func main() {
	var activeUsers int32
	var totalTimeSpent int64
	var wg sync.WaitGroup
	var sessionMap sync.Map
	var sessionMutex sync.Mutex

	// Create a channel to simulate user activity data stream
	userChannel := make(chan UserActivity, 100)

	// Number of goroutines to spawn
	numGoroutines := 5

	// Start multiple goroutines to process data from userChannel
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go processUserActivity(userChannel, &wg, &activeUsers, &totalTimeSpent, &sessionMap, &sessionMutex)
	}

	// Simulate user activity with random time spent
	for i := 0; i < 100; i++ {
		timeSpent := int64(1 + time.Now().UnixNano()%10)
		userChannel <- UserActivity{TimeSpent: timeSpent, StartTime: time.Now()}
	}

	// Close the channel to signal goroutines to stop processing
	close(userChannel)

	// Wait for all goroutines to finish
	wg.Wait()

	// Output the results
	fmt.Printf("Total Active Users: %d\n", activeUsers)
	fmt.Printf("Total Time Spent by Users: %d seconds\n", totalTimeSpent)
	averageSessionTime := calculateAverageSessionTime(&sessionMap)
	fmt.Printf("Average Session Time: %.2f seconds\n", averageSessionTime)
	peakActivityPeriod := findPeakActivityPeriod(&sessionMap)
	fmt.Printf("Peak Activity Period: %s - %s\n", peakActivityPeriod.StartTime.Format("15:04:05"), peakActivityPeriod.EndTime.Format("15:04:05"))
}

func processUserActivity(userChannel <-chan UserActivity, wg *sync.WaitGroup, activeUsers *int32, totalTimeSpent *int64, sessionMap *sync.Map, sessionMutex *sync.Mutex) {
	defer wg.Done()
	for activity := range userChannel {
		// Increment the shared active users counter atomically
		atomic.AddInt32(activeUsers, 1)

		// Add the time spent value to the total time spent atomically
		atomic.AddInt64(totalTimeSpent, activity.TimeSpent)

		// Update user session
		sessionMutex.Lock()
		session, ok := sessionMap.Load(activity.StartTime)
		if !ok {
			session = &Session{TotalTimeSpent: 0, StartTime: activity.StartTime}
		}
		s := session.(*Session)
		s.TotalTimeSpent += activity.TimeSpent
		sessionMap.Store(activity.StartTime, s)
		sessionMutex.Unlock()
	}
}

// calculateAverageSessionTime calculates the average session time.
func calculateAverageSessionTime(sessionMap *sync.Map) float64 {
	var totalSessions int
	var totalSessionTime int64
	sessionMap.Range(func(_, value interface{}) bool {
		session := value.(*Session)
		totalSessions++