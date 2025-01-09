package main

import (
	"fmt"
)

// Feedback struct to hold feedback details
type Feedback struct {
	UserID string `json:"user_id"`
	VideoID string `json:"video_id"`
	Rating int    `json:"rating"`
}

// FeedbackCollection struct to manage feedback entries
type FeedbackCollection struct {
	feedback map[string]map[string]Feedback
}

// NewFeedbackCollection creates a new feedback collection
func NewFeedbackCollection() *FeedbackCollection {
	return &FeedbackCollection{
		feedback: make(map[string]map[string]Feedback),
	}
}

// AddFeedback adds a new feedback entry
func (fc *FeedbackCollection) AddFeedback(userID, videoID string, rating int) {
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
	if videoFeedback, exists := fc.feedback[videoID]; exists {
		if feedback, userExists := videoFeedback[userID]; userExists {
			return feedback, true
		}
	}
	return Feedback{}, false
}

func main() {
	// Create a new feedback collection
	fc := NewFeedbackCollection()

	// Add some feedback entries
	fc.AddFeedback("user1", "video1", 5)
	fc.AddFeedback("user2", "video1", 4)
	fc.AddFeedback("user1", "video2", 3)

	// Retrieve feedback
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