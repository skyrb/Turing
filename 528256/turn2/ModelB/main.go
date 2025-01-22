package main

import (
	"fmt"
	"time"
)

type UserState struct {
	Timestamp time.Time
	Diffs     []string // Diff represents changes between states as strings
}

func main() {
	userStates := make([]UserState, 0)

	currentUsers := []string{"User1", "User2", "User3"}
	userStates = append(userStates, UserState{
		Timestamp: time.Now(),
		Diffs:     []string{"Initial state: User1, User2, User3"},
	})

	fmt.Println("Initial state:", currentUsers)

	// User logs in
	newUsers := append(currentUsers, "User4")
	diffs := calculateDiffs(currentUsers, newUsers)
	userStates = append(userStates, UserState{
		Timestamp: time.Now(),
		Diffs:     diffs,
	})
	currentUsers = newUsers

	fmt.Println("After User4 logs in:", currentUsers)

	// User logs out
	for i, user := range currentUsers {
		if user == "User3" {
			newUsers := append(currentUsers[:i], currentUsers[i+1:]...)
			diffs := calculateDiffs(currentUsers, newUsers)
			userStates = append(userStates, UserState{
				Timestamp: time.Now(),
				Diffs:     diffs,
			})
			currentUsers = newUsers
			break
		}
	}

	fmt.Println("After User3 logs out:", currentUsers)

	// Display all states with diffs
	fmt.Println("\nUser States with Diffs:")
	for i, state := range userStates {
		fmt.Printf("State %d: %v, Diffs: %v\n", i+1, state.Timestamp, state.Diffs)
	}

	// Example: Replay state changes to get the current state
	fmt.Println("\nReplaying state changes to get current state:")
	replayedUsers := []string{}
	for _, state := range userStates {
		for _, diff := range state.Diffs {
			replayedUsers = applyDiff(replayedUsers, diff)
		}
	}
	fmt.Println("Current state:", replayedUsers)
}

// calculateDiffs calculates the diffs between two slices of users
func calculateDiffs(oldUsers, newUsers []string) []string {
	diffs := make([]string, 0)
	// Implement your diff algorithm here, e.g., using LCS
	// For simplicity, we'll just append the changes as strings
	for _, user := range newUsers {
		if !contains(oldUsers, user) {
			diffs = append(diffs, fmt.Sprintf("%s logged in", user))
		}
	}
	for _, user := range oldUsers {
		if !contains(newUsers, user) {
			diffs = append(diffs, fmt.Sprintf("%s logged out", user))
		}
	}
	return diffs
}

// applyDiff applies a single diff to a slice of users
func applyDiff(users []string, diff string) []string {
	// Implement logic to apply the diff based on the diff string
	// For simplicity, we'll just parse the diff and update the slice
	if len(diff) == 0 {
		return users
	}

	var user string
	var action string
	_, err := fmt.Sscanf(diff, "%s %s", &user, &action)
	if err != nil {
		panic(err)
	}

	switch action {
	case "logged":
		if !contains(users, user) {
			users = append(users, user)
		}
	case "out":
		for i, u := range users {
			if u == user {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
	}
	return users
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
} 