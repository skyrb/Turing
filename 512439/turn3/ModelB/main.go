package main

import (
	"container/heap"
	"sync"
)

// Feedback struct to hold feedback details
type Feedback struct {
	UserID    string `json:"user_id"`
	VideoID   string `json:"video_id"`
	Rating    int    `json:"rating"`
	Priority int    `json:"priority"`
}

// FeedbackQueue represents a priority queue of Feedback
type FeedbackQueue []Feedback

func (fq FeedbackQueue) Len() int           { return len(fq) }
func (fq FeedbackQueue) Less(i, j int) bool { return fq[i].Priority > fq[j].Priority } // Change to < for low to high priority
func (fq FeedbackQueue) Swap(i, j int)      { fq[i], fq[j] = fq[j], fq[i] }

func (fq *FeedbackQueue) Push(x interface{}) {
	*fq = append(*fq, x.(Feedback))
}

func (fq *FeedbackQueue) Pop() interface{} {
	old := *fq
	n := len(old)
	x := old[n-1]
	*fq = old[0 : n-1]
	return x
}

// FeedbackCollection struct to manage feedback entries with priority queue
type FeedbackCollection struct {
	feedback       map[string]map[string]Feedback
	highPriorityQC FeedbackQueue
	mu             *sync.RWMutex
}

// NewFeedbackCollection creates a new feedback collection with concurrent access and priority queue
func NewFeedbackCollection() *FeedbackCollection {
	return &FeedbackCollection{
		feedback:       make(map[string]map[string]Feedback),
		highPriorityQC: FeedbackQueue{},
		mu:             &sync.RWMutex{},
	}
}

// AddFeedback adds a new feedback entry with priority
func (fc *FeedbackCollection) AddFeedback(userID, videoID string, rating, priority int) {
	feedback := Feedback{
		UserID:    userID,
		VideoID:   videoID,
		Rating:    rating,
		Priority: priority,
	}

	fc.mu.Lock()
	defer fc.mu.Unlock()

	if _, exists := fc.feedback[videoID]; !exists {
		fc.feedback[videoID] = make(map[string]Feedback)
	}
	fc.feedback[videoID][userID] = feedback

	// Add high priority feedback to the queue
	if priority > 0 {
		heap.Push(&fc.highPriorityQC, feedback)
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

// ProcessHighPriorityFeedback processes one high-priority feedback entry
func (fc *FeedbackCollection) ProcessHighPriorityFeedback() (Feedback, bool) {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	if len(fc.highPriorityQC) == 0 {
		return Feedback{}, false
	}
	feedback := heap.Pop(&fc.highPriorityQC).(Feedback)
	return feedback, true
}
    // ... rest of the code remains the same ...
