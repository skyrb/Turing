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
	users      map[int]UserProfile
	changes    chan ChangeRecord // Unbuffered channel to ensure all changes are logged
	mu         sync.RWMutex
	stop       chan struct{}
}

func NewLineageTracker() *LineageTracker {
	return &LineageTracker{
		users:    make(map[int]UserProfile),
		changes:  make(chan ChangeRecord),
		stop:     make(chan struct{}),
	}
}

func (lt *LineageTracker) Start() {
	go lt.logChanges()
}

func (lt *LineageTracker) logChanges() {
	for {
		select {
		case change := <-lt.changes:
			fmt.Printf("Change at %v: %s UserID %d - Old: %+v, New: %+v\n",
				change.Time, change.ChangeType, change.UserID, change.OldProfile, change.NewProfile)
		case <-lt.stop:
			log.Println("Lineage logger stopping.")
			return
		}
	}
}

func (lt *LineageTracker) AddUser(userID int, profile UserProfile) {
	lt.mu.Lock()
	defer lt.mu.Unlock()

	oldProfile, _ := lt.users[userID]
	lt.users[userID] = profile
	lt.changes <- ChangeRecord{
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
		lt.changes <- ChangeRecord{
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
	lt.changes <- ChangeRecord{
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