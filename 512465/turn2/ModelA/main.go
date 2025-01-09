package main

import (
	"fmt"
	"strings"
)

type Message struct {
	User    string
	Content string
}

type MessagingApp struct {
	Messages []Message
}

func (app *MessagingApp) AddMessage(user, content string) {
	app.Messages = append(app.Messages, Message{User: user, Content: content})
}

func (app *MessagingApp) SearchMessages(keywords []string) []Message {
	var results []Message
	keywordsLower := make([]string, len(keywords))
	for i, keyword := range keywords {
		keywordsLower[i] = strings.ToLower(keyword)
	}

	for _, message := range app.Messages {
		// Check both the User and Content fields for any of the keywords
		if containsAny(keywordsLower, strings.ToLower(message.Content)) ||
			containsAny(keywordsLower, strings.ToLower(message.User)) {
			results = append(results, message)
		}
	}
	return results
}

func containsAny(keys []string, s string) bool {
	for _, key := range keys {
		if strings.Contains(s, key) {
			return true
		}
	}
	return false
}

func main() {
	app := MessagingApp{}

	// Add some sample messages
	app.AddMessage("Alice", "Hi Bob! Let's meet at 3pm for lunch.")
	app.AddMessage("Bob", "Okay, Alice! See you then.")
	app.AddMessage("Charlie", "Just checking in. How are you?")
	app.AddMessage("Alice", "I'm fine, thanks for asking.")

	// Search for messages containing the keywords "Alice" and "Bob"
	keywords := []string{"Alice", "Bob"}
	results := app.SearchMessages(keywords)

	fmt.Printf("Messages containing '%v': \n", keywords)
	for _, result := range results {
		fmt.Printf("%s: %s\n", result.User, result.Content)
	}
}