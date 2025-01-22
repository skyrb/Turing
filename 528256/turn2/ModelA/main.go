package main

import (
    "fmt"
    "time"
)

// Defining operations as constants
const (
    Login  = "login"
    Logout = "logout"
)

// Change represents an incremental change to the user state
type Change struct {
    Timestamp time.Time
    User      string
    Action    string
}

// UserActivityTracker keeps track of user login/logout activity
type UserActivityTracker struct {
    BaseState []string
    Changes   []Change
}

func (uat *UserActivityTracker) AddChange(user string, action string) {
    change := Change{
        Timestamp: time.Now(),
        User:      user,
        Action:    action,
    }
    uat.Changes = append(uat.Changes, change)
    
    // Apply change to base state for demonstration
    if action == Login {
        uat.BaseState = append(uat.BaseState, user)
    } else if action == Logout {
        for i, u := range uat.BaseState {
            if u == user {
                uat.BaseState = append(uat.BaseState[:i], uat.BaseState[i+1:]...)
                break
            }
        }
    }
}

func (uat *UserActivityTracker) DisplayRecentActivity(duration time.Duration) {
    cutoff := time.Now().Add(-duration)
    fmt.Println("Activity in the last 24 hours:")
    for _, change := range uat.Changes {
        if change.Timestamp.After(cutoff) {
            fmt.Printf("%s: %s %s\n", change.Timestamp.Format(time.RFC3339), change.User, change.Action)
        }
    }
}

func main() {
    tracker := UserActivityTracker{
        BaseState: []string{"User1", "User2"},
        Changes:   []Change{},
    }

    // Log User3 logging in
    tracker.AddChange("User3", Login)
    // Log User2 logging out
    tracker.AddChange("User2", Logout)
    // Log User4 logging in
    tracker.AddChange("User4", Login)

    // Display current state
    fmt.Println("Current users:", tracker.BaseState)

    // Display user activity in the last 24 hours
    tracker.DisplayRecentActivity(24 * time.Hour)
}