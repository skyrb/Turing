package main

import (
	"fmt"
	"strconv"
)

type Post struct {
	ID     int
	Title  string
	Content string
}

// Observer interface for subscribing to blog post changes
type Observer interface {
	NotifyPostAdded(post Post)
	NotifyPostEdited(post Post)
	NotifyPostDeleted(post Post)
}

// Observable - The blogging platform struct that implements the Observer pattern
type Observable struct {
	observers []Observer
}

func (o *Observable) RegisterObserver(observer Observer) {
	o.observers = append(o.observers, observer)
}

func (o *Observable) NotifyAllObservers() {
	for _, observer := range o.observers {
		observer.NotifyPostAdded(posts[len(posts)-1]) // Notify all observers about the latest added post
	}
}

// LoggingObserver implements Observer interface for logging post changes
type LoggingObserver struct{}

func (l LoggingObserver) NotifyPostAdded(post Post) {
	fmt.Printf("LoggingObserver: Post added: '%s'\n", post.Title)
}

func (l LoggingObserver) NotifyPostEdited(post Post) {
	fmt.Printf("LoggingObserver: Post edited: '%s'\n", post.Title)
}

func (l LoggingObserver) NotifyPostDeleted(post Post) {
	fmt.Printf("LoggingObserver: Post deleted: '%s'\n", post.Title)
}

// AnalyticsObserver implements Observer interface for analytics post changes
type AnalyticsObserver struct{}

func (a AnalyticsObserver) NotifyPostAdded(post Post) {
	fmt.Printf("AnalyticsObserver: Post added to analytics: '%s'\n", post.Title)
}

func (a AnalyticsObserver) NotifyPostEdited(post Post) {
	fmt.Printf("AnalyticsObserver: Post edited in analytics: '%s'\n", post.Title)
}

func (a AnalyticsObserver) NotifyPostDeleted(post Post) {
	fmt.Printf("AnalyticsObserver: Post deleted from analytics: '%s'\n", post.Title)
}

var posts []Post
var observable Observable // Add the Observable struct instance

func main() {
	// Register observers
	observable.RegisterObserver(&LoggingObserver{})
	observable.RegisterObserver(&AnalyticsObserver{})

	menu()
}

func addPost() {
	// ... (existing code)

	observable.NotifyAllObservers() // Notify all observers about the added post
	menu()
}

// ... (Edit and Delete functions will also notify observers)

func editPost() {
	// ... (existing code)

	observable.NotifyAllObservers() // Notify all observers about the edited post
	menu()
}

func deletePost() {
	// ... (existing code)

	observable.NotifyAllObservers() // Notify all observers about the deleted post
	menu()
}