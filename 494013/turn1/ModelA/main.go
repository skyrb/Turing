package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type Notification struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Message string  `json:"message"`
}

var subscriptions = make(map[int]chan string, 100)
var subscriptionsMu sync.Mutex

func subscribeUser(userId int) (chan string, error) {
	subscriptionsMu.Lock()
	defer subscriptionsMu.Unlock()

	if _, ok := subscriptions[userId]; ok {
		return nil, fmt.Errorf("user %d already subscribed", userId)
	}

	ch := make(chan string)
	subscriptions[userId] = ch
	return ch, nil
}

func unsubscribeUser(userId int) {
	subscriptionsMu.Lock()
	defer subscriptionsMu.Unlock()

	delete(subscriptions, userId)
}

func processNotification(notification Notification) {
	subscriptionsMu.Lock()
	defer subscriptionsMu.Unlock()

	for _, ch := range subscriptions {
		select {
		case ch <- fmt.Sprintf("Notification %d: %s - %s", notification.ID, notification.Title, notification.Message):
			default:
				// If the channel is full or closed, don't do anything
		}
	}
}

func main() {
	user1 := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	user2 := User{ID: 2, Name: "Bob", Email: "bob@example.com"}

	// Subscribe users
	ch1, err := subscribeUser(user1.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch2, err := subscribeUser(user2.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		// Simulate receiving notifications and processing them
		for i := 1; i <= 5; i++ {
			notification := Notification{ID: i, Title: fmt.Sprintf("Event %d", i), Message: fmt.Sprintf("This is an event for Event %d", i)}
			processNotification(notification)
			fmt.Printf("Processed notification %d\n", notification.ID)
			<-time.After(time.Second * 2)
		}
	}()

	// Simulate user reading notifications
	for {
		select {
		case msg := <-ch1:
			fmt.Println("Alice received:", msg)
		case msg := <-ch2:
			fmt.Println("Bob received:", msg)
		}
	}
}