package main

import (
	"fmt"
	"time"
)

type UserState struct {
	Timestamp time.Time
	Users     []string
	Changes   []string
}

type ActivityState struct {
	Timestamp time.Time
	Data      []string // Replace []string with your actual data type for each activity
	Changes   []string
}

type HistoryEntry struct {
	Timestamp time.Time
	Activity  string
	Changes   []string
}

var activityStates map[string]ActivityState
var history []HistoryEntry

func init() {
	activityStates = make(map[string]ActivityState)
	history = make([]HistoryEntry, 0)
}

func main() {
	// Initialize states for different activities
	activityStates["user_activity"] = ActivityState{
		Timestamp: time.Now(),
		Users:     []string{"User1", "User2", "User3"},
		Changes:   []string{"Initial user activity state"},
	}

	activityStates["file_uploads"] = ActivityState{
		Timestamp: time.Now(),
		Data:      []string{"file1.txt", "file2.png"},
		Changes:   []string{"Initial file upload state"},
	}

	fmt.Println("Initial states:")
	fmt.Println("User activity:", activityStates["user_activity"].Users)
	fmt.Println("File uploads:", activityStates["file_uploads"].Data)

	// User logs in (affects user_activity)
	newUsers := append(activityStates["user_activity"].Users, "User4")
	updateActivityState("user_activity", newUsers, []string{"User4 logged in"})

	// File uploaded (affects file_uploads)
	newFiles := append(activityStates["file_uploads"].Data, "file3.doc")
	updateActivityState("file_uploads", newFiles, []string{"file3.doc uploaded"})

	// User logs out (affects user_activity)
	for i, user := range activityStates["user_activity"].Users {
		if user == "User3" {
			newUsers := append(activityStates["user_activity"].Users[:i], activityStates["user_activity"].Users[i+1:]...)
			updateActivityState("user_activity", newUsers, []string{"User3 logged out"})
			break
		}
	}

	// Display history
	fmt.Println("\nHistory:")
	for _, entry := range history {
		fmt.Printf("%v: %s: %v\n", entry.Timestamp, entry.Activity, entry.Changes)
	}
}

func updateActivityState(activity string, newData []string, changes []string) {
	oldState, ok := activityStates[activity]
	if !ok {
		panic("Activity not found")
	}

	newState := ActivityState{
		Timestamp: time.Now(),
		Data:      newData,
		Changes:   changes,
	}

	activityStates[activity] = newState

	history = append(history, HistoryEntry{
		Timestamp: time.Now(),
		Activity:  activity,
		Changes:   changes,
	})
}