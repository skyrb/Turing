
package main  
import (  
    "fmt"
    "sync"
)  

type Feedback struct {
    UserID   int
    VideoID  int
    Rating   int
}

type FeedbackCollection struct {
    sync.Mutex
    FeedbackMap map[int]map[int]Feedback
}

// NewFeedbackCollection creates a new FeedbackCollection
func NewFeedbackCollection() *FeedbackCollection {
    return &FeedbackCollection{
        FeedbackMap: make(map[int]map[int]Feedback),
    }
}

// AddFeedback adds a new feedback to the collection
func (fc *FeedbackCollection) AddFeedback(feedback Feedback) {
    fc.Lock()
    defer fc.Unlock()

    userID := feedback.UserID
    videoID := feedback.VideoID

    // Check if the userID already exists in the FeedbackMap
    if _, ok := fc.FeedbackMap[userID]; !ok {
        // If userID doesn't exist, initialize a new map for their feedback
        fc.FeedbackMap[userID] = make(map[int]Feedback)
    }

    // Add the feedback for the specified userID and videoID
    fc.FeedbackMap[userID][videoID] = feedback
}

// GetFeedback retrieves feedback for a given userID and videoID
func (fc *FeedbackCollection) GetFeedback(userID int, videoID int) (Feedback, bool) {
    fc.Lock()
    defer fc.Unlock()

    userFeedback, ok := fc.FeedbackMap[userID]
    if !ok {
        return Feedback{}, false // UserID not found
    }

    feedback, found := userFeedback[videoID]
    return feedback, found
}

// GetFeedbackForUser retrieves all feedback for a given userID
func (fc *FeedbackCollection) GetFeedbackForUser(userID int) ([]Feedback, bool) {
    fc.Lock()
    defer fc.Unlock()

    userFeedback, ok := fc.FeedbackMap[userID]
    if !ok {
        return nil, false // UserID not found
    }

    var feedbacks []Feedback
    for _, feedback := range userFeedback {
        feedbacks = append(feedbacks, feedback)
    }
    return feedbacks, true
}

// RemoveFeedback removes feedback for a given userID and videoID
func (fc *FeedbackCollection) RemoveFeedback(userID int, videoID int) {
    fc.Lock()
    defer fc.Unlock()

    userFeedback, ok := fc.FeedbackMap[userID]
    if !ok {
        return // UserID not found
    }

    delete(userFeedback, videoID)

    // If the user has no more feedback, remove their entry from the outer map
    if len(userFeedback) == 0 {
        delete(fc.FeedbackMap, userID)
    }
}

func main() {
    // Create a new feedback collection
    feedbackCollection := NewFeedbackCollection()
    
    // Add some feedback entries
    feedbackCollection.AddFeedback(Feedback{UserID: 1, VideoID: 100, Rating: 4})
    feedbackCollection.AddFeedback(Feedback{UserID: 2, VideoID: 101, Rating: 3})
    feedbackCollection.AddFeedback(Feedback{UserID: 1, VideoID: 102, Rating: 5})

    // Get feedback for userID 1 and videoID 100
    feedback, found := feedbackCollection.GetFeedback(1, 100)
    if found {
        fmt.Println("Feedback for User 1, Video 100:", feedback)
    } else {
        fmt.Println("Feedback not found.")
    }

    // Get feedback for userID 2
    feedbacks, found := feedbackCollection.GetFeedbackForUser(2)
    if found {
        fmt.Println("Feedback for User 2:")
        for _, feedback := range feedbacks {
            fmt.Println(feedback)
        }
    } else {
        fmt.Println("Feedback not found.")
    }

    // Remove feedback for userID 1, videoID 100
    feedbackCollection.RemoveFeedback(1, 100)