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

func (app *MessagingApp) SearchMessages(keyword string) []Message {
	var results []Message
	keyword = strings.ToLower(keyword)

	for _, message := range app.Messages {
		if strings.Contains(strings.ToLower(message.Content), keyword) {
			results = append(results, message)
		}
	}
	return results
}

func main() {
	app := MessagingApp{}

	// Add some sample messages
	app.AddMessage("Alice", "Hi Bob! Let's meet at 3pm for lunch.")
	app.AddMessage("Bob", "Okay, Alice! See you then.")
	app.AddMessage("Charlie", "Just checking in. How are you?")
	app.AddMessage("Alice", "I'm fine, thanks for asking.")

	// Search for messages containing the keyword "meeting"
	keyword := "Alice"
	results := app.SearchMessages(keyword)

	fmt.Printf("Messages containing '%s': \n", keyword)
	for _, result := range results {
		fmt.Printf("%s: %s\n", result.User, result.Content)
	}
}