package main

import (
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	Type    string
	Data    interface{}
	Priority int // Higher numbers indicate higher priority
}

var (
	// Map of user ID to a slice of their notification channels.
	subscriptions       map[int][]chan Notification
	subscriptionsMutex  sync.Mutex

	// Map of group IDs to a slice of user IDs.
	groups               map[string][]int
	groupsMutex          sync.Mutex

	// Priority queue to manage notifications based on priority.
	priorityQueue []Notification
	priorityQueueMutex sync.Mutex
)

func init() {
	subscriptions = make(map[int][]chan Notification)
	groups = make(map[string][]int)
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

func addUserToGroup(groupID string, userID int) {
	groupsMutex.Lock()
	defer groupsMutex.Unlock()

	groups[groupID] = append(groups[groupID], userID)
}

func removeUserFromGroup(groupID string, userID int) {
	groupsMutex.Lock()
	defer groupsMutex.Unlock()

	// Find the index of the user ID in the slice for the group
	index := -1
	for i, u := range groups[groupID] {
		if u == userID {
			index = i
			break
		}
	}
	if index != -1 {
		// Remove the user ID from the slice
		groups[groupID] = append(groups[groupID][:index], groups[groupID][index+1:]...)
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

func notifyGroup(groupID string, notification Notification) {
	groupsMutex.Lock()
	defer groupsMutex.Unlock()

	// Find all the users in the group
	for _, userID := range groups[groupID] {
		notifyUser(userID, notification)
	}
}

func sendPriorityNotifications() {
	priorityQueueMutex.Lock()
	defer priorityQueueMutex.Unlock()

	for len(priorityQueue) > 0 {
		highestPriority := priorityQueue[0]
		priorityQueue = priorityQueue[1:]

		// Send the highest priority notification
		notifyUser(highestPriority.Type, highestPriority)
	}
}

func sendNotification(notification Notification) {
	priorityQueueMutex.Lock()
	defer priorityQueueMutex.Unlock()

	// Add the notification to the priority queue
	for i := len(priorityQueue) - 1; i >= 0; i-- {
		if priorityQueue[i].Priority < notification.Priority {
			priorityQueue = append(priorityQueue, Notification{})
			copy(priorityQueue[i+1:], priorityQueue[i:])
			priorityQueue[i] = notification
			return
		}
	}
	priorityQueue = append(priorityQueue, notification)

	// Send the notification
	go sendPriorityNotifications()
}

func listenForNotifications(userID int, notifications chan Notification) {
	for notification := range notifications {
		fmt.Printf("User %d received notification: %s %v\n", userID, notification.Type, notification.Data)
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

	// Add users to groups
	addUserToGroup("group1", userID1)
	addUserToGroup("group1", userID2)
	addUserToGroup("group2", userID3)

	// Start listening for notifications for each user in separate goroutines
	go listenForNotifications(userID1, notifications1)
	go listenForNotifications(userID2, notifications2)
	go listenForNotifications(userID3, notifications3)

	// Simulate sending notifications
	go func() {
		for {
			time.Sleep(2 * time.Second)
			sendNotification(Notification{Type: "Update", Data: "New data available!", Priority: 1})

			time.Sleep(3 * time.Second)
			sendNotification(Notification{Type: "Reminder", Data: "Don't forget to do something!", Priority: 2})

			time.Sleep(4 * time.Second)
			notifyGroup("group1", Notification{Type: "Broadcast", Data: "Group 1 announcement!", Priority: 3})
		}
	}()

	// Keep the main function running
	select {}
}