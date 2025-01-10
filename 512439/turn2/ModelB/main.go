package main  
import (  
    "fmt"
    "sync"
)

const (  
    initialMapSize = 1000000 // Initial size of the map (adjust this based on expected peak traffic)
    minMapSize     = 1000    // Minimum size of the map after resizing down
    loadFactor    = 0.7    // Load factor to trigger map resizing (60% of capacity)
)

type Feedback struct {  
    UserID string `json:"user_id"`
    VideoID string `json:"video_id"`
    Rating int `json:"rating"`
}

type FeedbackCollection struct {  
    feedback map[string]map[string]Feedback
    mu       sync.Mutex
}

func NewFeedbackCollection() *FeedbackCollection {  
    return &FeedbackCollection{  
        feedback: make(map[string]map[string]Feedback, initialMapSize),
    }
}

func (fc *FeedbackCollection) AddFeedback(userID, videoID string, rating int) {  
    fc.mu.Lock()
    defer fc.mu.Unlock()

    if _, exists := fc.feedback[videoID]; !exists {  
        fc.feedback[videoID] = make(map[string]Feedback)
    }
    fc.feedback[videoID][userID] = Feedback{  
        UserID:  userID,
        VideoID: videoID,
        Rating:  rating,
    }

    fc.resizeIfNeeded()
}

func (fc *FeedbackCollection) GetFeedback(userID, videoID string) (Feedback, bool) {  
    fc.mu.Lock()
    defer fc.mu.Unlock()

    if videoFeedback, exists := fc.feedback[videoID]; exists {  
        if feedback, userExists := videoFeedback[userID]; userExists {  
            return feedback, true
        }
    }
    return Feedback{}, false
}

func (fc *FeedbackCollection) resizeIfNeeded() {  
    if len(fc.feedback) > cap(fc.feedback)*loadFactor {  
        newFeedback := make(map[string]map[string]Feedback, len(fc.feedback)*2) // Double the size
        for videoID, userFeedback := range fc.feedback {  
            newFeedback[videoID] = make(map[string]Feedback, len(userFeedback))
            for userID, feedback := range userFeedback {  
                newFeedback[videoID][userID] = feedback
            }
        }
        fc.feedback = newFeedback
    } else if len(fc.feedback) < cap(fc.feedback)/2 && cap(fc.feedback) > minMapSize {  
        // Downsize the map if the load factor is too low and it's still larger than the minimum size
        newFeedback := make(map[string]map[string]Feedback, cap(fc.feedback)/2)
        for videoID, userFeedback := range fc.feedback {  
            newFeedback[videoID] = make(map[string]Feedback, len(userFeedback))
            for userID, feedback := range userFeedback {  
                newFeedback[videoID][userID] = feedback
            }
        }
        fc.feedback = newFeedback
    }
}

func main() {  
    fc := NewFeedbackCollection()
    // Adding feedback entries (can be done concurrently from multiple goroutines)
    // ...
}