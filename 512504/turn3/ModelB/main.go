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
    OpType      string // "Added", "Updated", "Deleted"
}

type LineageTracker struct {
    users      map[int]UserProfile
    addCh      chan UserProfile
    updateCh   chan UserProfile
    deleteCh   chan int
    changeCh   chan ChangeRecord
    mu         sync.RWMutex
    stop       chan struct{}
}

func NewLineageTracker() *LineageTracker {  
    return &LineageTracker{
        users:     make(map[int]UserProfile),
        addCh:     make(chan UserProfile),
        updateCh:   make(chan UserProfile),
        deleteCh:   make(chan int),
        changeCh:   make(chan ChangeRecord),
        stop:      make(chan struct{}),
    }
}

func (lt *LineageTracker) Start() {
    go lt.addHandler()
    go lt.updateHandler()
    go lt.deleteHandler()
    go lt.logChanges()
}

func (lt *LineageTracker) addHandler() {
    for {
        select {
        case user := <-lt.addCh:
            lt.mu.Lock()
            lt.users[user.Age] = user
            lt.changeCh <- ChangeRecord{
                Time:        time.Now(),
                UserID:      user.Age,
                OldProfile:  UserProfile{},
                NewProfile:  user,
                OpType:      "Added",
            }
            lt.mu.Unlock()
        case <-lt.stop:
            return
        }
    }
}

func (lt *LineageTracker) updateHandler() {
    for {
        select {
        case user := <-lt.updateCh:
            lt.mu.Lock()
            oldProfile, exists := lt.users[user.Age]
            if !exists {
                log.Printf("UserID %d not found for update\n", user.Age)
            } else {
                lt.users[user.Age] = user
                lt.changeCh <- ChangeRecord{
                    Time:        time.Now(),
                    UserID:      user.Age,
                    OldProfile:  oldProfile,
                    NewProfile:  user,
                    OpType:      "Updated",
                }
            }
            lt.mu.Unlock()
        case <-lt.stop:
            return
        }
    }
}

func (lt *LineageTracker) deleteHandler() {
    for {
        select {
        case userID := <-lt.deleteCh:
            lt.mu.Lock()
            oldProfile, exists := lt.users[userID]
            if !exists {
                log.Printf("UserID %d not found for deletion\n", userID)
            } else {
                delete(lt.users, userID)
                lt.changeCh <- ChangeRecord{
                    Time:        time.Now(),
                    UserID:      userID,
                    OldProfile:  oldProfile,