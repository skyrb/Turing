package main

import (
	"errors"
	"fmt"
)

// UserActivityTracker maintains the list of active users and version history.
type UserActivityTracker struct {
	// users represents the current state of logged-in users.
	users []string
	// history maintains a history of user states for version control.
	history [][]string
}

// NewUserActivityTracker initializes a new UserActivityTracker.
func NewUserActivityTracker() *UserActivityTracker {
	return &UserActivityTracker{
		users:   []string{},
		history: [][]string{},
	}
}

// AddUser logs a user in by adding them to the users slice.
func (uat *UserActivityTracker) AddUser(user string) {
	uat.users = append(uat.users, user)
	uat.recordState() // record the new state
}

// RemoveUser logs a user out by removing them from the users slice.
func (uat *UserActivityTracker) RemoveUser(user string) error {
	index := -1
	for i, u := range uat.users {
		if u == user {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("user not found")
	}
	uat.users = append(uat.users[:index], uat.users[index+1:]...)
	uat.recordState() // record the new state
	return nil
}

// recordState captures the current state of the users slice.
func (uat *UserActivityTracker) recordState() {
	// Create a deep copy of the current users slice and add it to history.
	copyOfCurrentState := make([]string, len(uat.users))
	copy(copyOfCurrentState, uat.users)
	uat.history = append(uat.history, copyOfCurrentState)
}

// GetVersion returns the state of users at a specific version index.
func (uat *UserActivityTracker) GetVersion(version int) ([]string, error) {
	if version < 0 || version >= len(uat.history) {
		return nil, errors.New("version index out of range")
	}
	return uat.history[version], nil
}

func main() {
	tracker := NewUserActivityTracker()
	tracker.AddUser("User1")
	tracker.AddUser("User2")
	tracker.AddUser("User3")
	fmt.Println("Current State:", tracker.users)

	err := tracker.RemoveUser("User2")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current State after removal:", tracker.users)

	// Print the history of states
	for i, state := range tracker.history {
		fmt.Printf("Version %d: %v\n", i, state)
	}

	// Fetch a specific version
	if version, err := tracker.GetVersion(1); err == nil {
		fmt.Println("Version 1:", version)
	} else {
		fmt.Println("Error getting version:", err)
	}
}