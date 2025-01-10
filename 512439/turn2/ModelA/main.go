package main

import (
	"fmt"
	"sync"
)

// Feedback struct to hold feedback details
type Feedback struct {
	UserID string `json:"user_id"`
	VideoID string `json:"video_id"`
	Rating int    `json:"rating"`
}

// FeedbackCollection struct to manage feedback entries
type FeedbackCollection struct {
	feedback  map[string]map[string]Feedback
	mu        *sync.RWMutex
}

// NewFeedbackCollection creates a new feedback collection with concurrent access
func NewFeedbackCollection() *FeedbackCollection {
	return &FeedbackCollection{
		feedback: make(map[string]map[string]Feedback),
		mu:       &sync.RWMutex{},
	}
}

// AddFeedback adds a new feedback entry
func (fc *FeedbackCollection) AddFeedback(userID, videoID string, rating int) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	if _, exists := fc.feedback[videoID]; !exists {
		fc.feedback[videoID] = make(map[string]Feedback)
	}
	fc.feedback[videoID][userID] = Feedback{
		UserID: userID,
		VideoID: videoID,
		Rating: rating,
	}
}

// GetFeedback retrieves feedback for a specific user and video
func (fc *FeedbackCollection) GetFeedback(userID, videoID string) (Feedback, bool) {
	fc.mu.RLock()
	defer fc.mu.RUnlock()

	if videoFeedback, exists := fc.feedback[videoID]; exists {
		if feedback, userExists := videoFeedback[userID]; userExists {
			return feedback, true
		}
	}
	return Feedback{}, false
}

// AddFeedbackSimulate simulates adding feedback concurrently
func AddFeedbackSimulate(fc *FeedbackCollection, numFeedback int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numFeedback; i++ {
		userID := fmt.Sprintf("user%d", i%100)
		videoID := fmt.Sprintf("video%d", i%10)
		rating := (i % 5) + 1 // Random rating between 1 and 5
		fc.AddFeedback(userID, videoID, rating)
	}
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
		go AddFeedbackSimulate(fc, numFeedbackPerGoroutine, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Example of retrieving feedback
	feedback, exists := fc.GetFeedback("user1", "video1")
	if exists {
		fmt.Printf("Feedback from %s for %s: %d\n", feedback.UserID, feedback.VideoID, feedback.Rating)
	} else {
		fmt.Println("Feedback not found.")
	}

	feedback, exists = fc.GetFeedback("user3", "video1")
	if exists {
		fmt.Printf("Feedback from %s for %s: %d\n", feedback.UserID, feedback.VideoID, feedback.Rating)
	} else {
		fmt.Println("Feedback not found.")
	}
}