package main

import (
	"fmt"
	"time"
)

// Activity interface defines the methods for version control
type Activity interface {
	CalculateDiffs(oldState, newState []string) []string
	ApplyDiff(state []string, diff string) []string
	GetName() string
}

// Diff struct represents a change in the activity state
type Diff struct {
	Timestamp time.Time
	Changes   []string
}

// ActivityManager manages multiple activities
type ActivityManager struct {
	activities map[string][]Diff
}

// NewActivityManager initializes a new ActivityManager
func NewActivityManager() *ActivityManager {
	return &ActivityManager{
		activities: make(map[string][]Diff),
	}
}

// RecordState records a new state for an activity
func (am *ActivityManager) RecordState(activity Activity, oldState, newState []string) {
	diffs := activity.CalculateDiffs(oldState, newState)
	if len(diffs) > 0 {
		am.activities[activity.GetName()] = append(am.activities[activity.GetName()], Diff{
			Timestamp: time.Now(),
			Changes:   diffs,
		})
	}
}

// GetHistory retrieves the history of changes for an activity
func (am *ActivityManager) GetHistory(activity Activity) []Diff {
	return am.activities[activity.GetName()]
}

// UserLoginActivity implements the Activity interface for user logins
type UserLoginActivity struct{}

func (ula *UserLoginActivity) CalculateDiffs(oldState, newState []string) []string {
	diffs := make([]string, 0)
	for _, user := range newState {
		if !contains(oldState, user) {
			diffs = append(diffs, fmt.Sprintf("%s logged in", user))
		}
	}
	for _, user := range oldState {
		if !contains(newState, user) {
			diffs = append(diffs, fmt.Sprintf("%s logged out", user))
		}
	}
	return diffs
}

func (ula *UserLoginActivity) ApplyDiff(state []string, diff string) []string {
	var user, action string
	_, err := fmt.Sscanf(diff, "%s %s", &user, &action)
	if err != nil {
		panic(err)
	}

	switch action {
	case "logged":
		if !contains(state, user) {
			state = append(state, user)
		}
	case "out":
		for i, u := range state {
			if u == user {
				state = append(state[:i], state[i+1:]...)
				break
			}
		}
	}
	return state
}

func (ula *UserLoginActivity) GetName() string {
	return "UserLogin"
}

// Helper function to check if a slice contains an element
func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func main() {
	manager := NewActivityManager()

	userActivity := &UserLoginActivity{}

	// Initial State
	currentUsers := []string{"User1", "User2", "User3"}
	manager.RecordState(userActivity, []string{}, currentUsers)

	// User logs in
	newUsers := append(currentUsers, "User4")
	manager.RecordState(userActivity, currentUsers, newUsers)
	currentUsers = newUsers

	// User logs out
	for i, user := range currentUsers {
		if user == "User3" {
			newUsers := append(currentUsers[:i], currentUsers[i+1:]...)
			manager.RecordState(userActivity, currentUsers, newUsers)
			currentUsers = newUsers
			break
		}
	}

	// Retrieve and display user login history
	fmt.Println("User Login History:")
	for _, diff := range manager.GetHistory(userActivity) {
		fmt.Printf("%v: %v\n", diff.Timestamp, diff.Changes)
	}
}