package main

import (
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	Type string
	Data interface{}
}

var (
	// Map of user ID to a slice of their notification channels.
	subscriptions map[int][]chan Notification
	subscriptionsMutex sync.Mutex
)

func init() {
	subscriptions = make(map[int][]chan Notification)
}

func subscribe(userID int, notifications chan Notification) {
	subscriptionsMutex.Lock()
	defer subscriptionsMutex.Unlock()

	subscriptions[userID] = append(subscriptions[userID], notifications)
}

func unsubscribe(userID int, notifications chan Notification) {
	subscriptionsMutex.Lock()
	defer subscriptionsMutex.Unlock()

	// Find the index of the notification channel in the slice for the user
	index := -1
	for i, n := range subscriptions[userID] {
		if n == notifications {
			index = i
			break
		}
	}
	if index != -1 {
		// Remove the channel from the slice
		subscriptions[userID] = append(subscriptions[userID][:index], subscriptions[userID][index+1:]...)
	}
}

func notifyUser(userID int, notification Notification) {
	subscriptionsMutex.Lock()
	defer subscriptionsMutex.Unlock()
	
	// Find all the channels the user is subscribed to
	for _, ch := range subscriptions[userID] {
		// Send the notification to the channel asynchronously
		go func(ch chan Notification) {
			ch <- notification
		}(ch)
	}
}

func main() {
	// Example usage:
	userID1 := 1
	userID2 := 2
	userID3 := 3

	// Create channels for each user to receive notifications
	notifications1 := make(chan Notification)
	notifications2 := make(chan Notification)
	notifications3 := make(chan Notification)

	// Subscribe users
	subscribe(userID1, notifications1)
	subscribe(userID2, notifications2)
	subscribe(userID3, notifications3)
	subscribe(userID1, notifications3)

	// Unsubscribed users
	unsubscribe(userID2, notifications2)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			// Simulate sending a notification to user 1
			notifyUser(userID1, Notification{Type: "Update", Data: "New data available!"})

			time.Sleep(time.Second * 3)
			// Simulate sending a notification to user 3
			notifyUser(userID3, Notification{Type: "Reminder", Data: "Don't forget to do something!"})
		}
	}()

	// Receive notifications for user 1
	for notification := range notifications1 {
		fmt.Println("User 1 received notification:", notification.Type, notification.Data)
	}

	// Receive notifications for user 3
	for notification := range notifications3 {
		fmt.Println("User 3 received notification:", notification.Type, notification.Data)
	}
} 