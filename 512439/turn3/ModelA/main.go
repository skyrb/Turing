package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Feedback struct to hold feedback details
type Feedback struct {
	UserID string   `json:"user_id"`
	VideoID string   `json:"video_id"`
	Rating int      `json:"rating"`
	Priority int     `json:"priority"` // 0 is low, 1 is high
}

// FeedbackCollection struct to manage feedback entries
type FeedbackCollection struct {
	highPriority    chan Feedback
	lowPriority    chan Feedback
	processHighMutex sync.Mutex
	stop            bool
}

// NewFeedbackCollection creates a new feedback collection with concurrent access
func NewFeedbackCollection() *FeedbackCollection {
	return &FeedbackCollection{
		highPriority:    make(chan Feedback),
		lowPriority:    make(chan Feedback),
		processHighMutex: sync.Mutex{},
		stop:            false,
	}
}

// AddFeedback adds a new feedback entry with a priority level
func (fc *FeedbackCollection) AddFeedback(userID, videoID string, rating int, priority int) {
	feedback := Feedback{
		UserID: userID,
		VideoID: videoID,
		Rating: rating,
		Priority: priority,
	}
	if priority == 1 {
		fc.highPriority <- feedback
	} else {
		fc.lowPriority <- feedback
	}
}

// ProcessFeedback consumes feedback from both queues, prioritizing high-priority entries
func (fc *FeedbackCollection) ProcessFeedback(wsg *sync.WaitGroup) {
	defer wsg.Done()

	for !fc.stop {
		// Process high-priority queue first
		select {
		case feedback, ok := <-fc.highPriority:
			if !ok {
				continue
			}
			fmt.Printf("Processing high-priority feedback from %s for %s: %d\n", feedback.UserID, feedback.VideoID, feedback.Rating)

		default:
			// If high-priority queue is empty, process low-priority queue
			select {
			case feedback, ok := <-fc.lowPriority:
				if !ok {
					continue
				}
				fmt.Printf("Processing low-priority feedback from %s for %s: %d\n", feedback.UserID, feedback.VideoID, feedback.Rating)

			default:
				// If both queues are empty, briefly sleep to avoid unnecessary processing
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

// Stop stops processing feedback
func (fc *FeedbackCollection) Stop() {
	fc.stop = true
	close(fc.highPriority)
	close(fc.lowPriority)
}

func main() {
	// Create a new feedback collection
	fc := NewFeedbackCollection()

	// Simulate adding feedback concurrently
	var wg sync.WaitGroup
	numGoroutines := 10
	numFeedbackPerGoroutine := 10000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numFeedbackPerGoroutine; j++ {
				userID := fmt.Sprintf("user%d", j%100)
				videoID := fmt.Sprintf("video%d", j%10)
				rating := (j % 5) + 1 // Random rating between 1 and 5
				priority := atomic.AddInt32(&j, 1) % 2 // Random priority of 0 or 1
				fc.AddFeedback(userID, videoID, rating, int(priority))
			}
		}()
	}

	// Start processing feedback
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go fc.ProcessFeedback(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Stop processing feedback
	fc.Stop()
}