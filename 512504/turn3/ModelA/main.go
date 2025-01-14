package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	Name    string
	Age     int
	Location string
}

type ChangeRecord struct {
	Time        time.Time
	UserID      int
	OldProfile  UserProfile
	NewProfile  UserProfile
	ChangeType  string // e.g., "Updated", "Added", "Deleted"
}

type LineageTracker struct {
	users       map[int]UserProfile
	addChanges  chan ChangeRecord
	updateChanges chan ChangeRecord
	deleteChanges chan ChangeRecord
	mu           sync.RWMutex
	stop         chan struct{}
}

func NewLineageTracker() *LineageTracker {
	return &LineageTracker{
		users:       make(map[int]UserProfile),
		addChanges:  make(chan ChangeRecord),
		updateChanges: make(chan ChangeRecord),
		deleteChanges: make(chan ChangeRecord),
		stop:         make(chan struct{}),
	}
}

func (lt *LineageTracker) Start() {
	go lt.logAddChanges()
	go lt.logUpdateChanges()
	go lt.logDeleteChanges()
}

func (lt *LineageTracker) logAddChanges() {
	for {
		select {
		case change := <-lt.addChanges:
			fmt.Printf("Add at %v: UserID %d - New: %+v\n", change.Time, change.UserID, change.NewProfile)
		case <-lt.stop:
			log.Println("Add lineage logger stopping.")
			return
		}
	}
}

func (lt *LineageTracker) logUpdateChanges() {
	for {
		select {
		case change := <-lt.updateChanges:
			fmt.Printf("Update at %v: UserID %d - Old: %+v, New: %+v\n",
				change.Time, change.UserID, change.OldProfile, change.NewProfile)
		case <-lt.stop:
			log.Println("Update lineage logger stopping.")
			return
		}
	}
}

func (lt *LineageTracker) logDeleteChanges() {
	for {
		select {
		case change := <-lt.deleteChanges:
			fmt.Printf("Delete at %v: UserID %d - Old: %+v\n", change.Time, change.UserID, change.OldProfile)
		case <-lt.stop:
			log.Println("Delete lineage logger stopping.")
			return
		}
	}
}

func (lt *LineageTracker) AddUser(userID int, profile UserProfile) {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	oldProfile, _ := lt.users[userID]
	lt.users[userID] = profile
	lt.addChanges <- ChangeRecord{
		Time:        time.Now(),
		UserID:      userID,
		OldProfile:  oldProfile,
		NewProfile:  profile,
		ChangeType:  "Added",
	}
}

func (lt *LineageTracker) UpdateUser(userID int, profile UserProfile) {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	oldProfile, exists := lt.users[userID]
	if !exists {
		log.Printf("UserID %d not found\n", userID)
		return
	}

	if oldProfile.Location != profile.Location { // Log only if the Location changes
		lt.users[userID] = profile
		lt.updateChanges <- ChangeRecord{
			Time:        time.Now(),
			UserID:      userID,
			OldProfile:  oldProfile,
			NewProfile:  profile,
			ChangeType:  "Updated",
		}
	}
}

func (lt *LineageTracker) DeleteUser(userID int) {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	oldProfile, exists := lt.users[userID]
	if !exists {
		log.Printf("UserID %d not found\n", userID)
		return
	}

	delete(lt.users, userID)
	lt.deleteChanges <- ChangeRecord{
		Time:        time.Now(),
		UserID:      userID,
		OldProfile:  oldProfile,
		NewProfile:  UserProfile{},
		ChangeType:  "Deleted",
	}
}

func (lt *LineageTracker) Stop() {
	close(lt.stop)
}

func main() {
	tracker := NewLineageTracker()
	tracker.Start()

	// Example of concurrent user modifications
	go func() {
		tracker.AddUser(1, UserProfile{Name: "Alice", Age: 25, Location: "New York"})
		time.Sleep(time.Second)
		tracker.UpdateUser(1, UserProfile{Name: "Alice", Age: 26, Location: "San Francisco"}) // Location change
		tracker.UpdateUser(1, UserProfile{Name: "Alice", Age: 26, Location: "San Francisco"}) // No Location change
		tracker.DeleteUser(1)
	}()

	go func() {
		tracker.AddUser(2, UserProfile{Name: "Bob", Age: 30, Location: "Los Angeles"})
		time.Sleep(time.Second)
		tracker.UpdateUser(2, UserProfile{Name: "Robert", Age: 30, Location: "Los Angeles"})
	}()

	// Wait for the processes to complete
	time.Sleep(3 * time.Second)

	tracker.Stop()
}