package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Message struct {
	User    string
	Content string
	Time    string // Example time format: "2023-10-04T10:30:00Z"
}

var messages []Message

func sendMessage(user, content string) {
	time := "2023-10-04T10:30:00Z" // For demonstration purposes
	message := Message{User: user, Content: content, Time: time}
	messages = append(messages, message)
	fmt.Printf("Message from %s: %s\n", user, content)
}

func searchMessages(keyword string) []Message {
	var filteredMessages []Message
	for _, message := range messages {
		if strings.Contains(message.Content, keyword) {
			filteredMessages = append(filteredMessages, message)
		}
	}
	return filteredMessages
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter username: ")
	user := scanner.Scan() && scanner.Text()
	if !user {
		fmt.Println("Invalid username")
		return
	}

	for {
		fmt.Print("Enter a message or 'search' to search: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "search" {
			fmt.Print("Enter keyword to search for: ")
			scanner.Scan()
			keyword := scanner.Text()
			results := searchMessages(keyword)

			if len(results) == 0 {
				fmt.Printf("No messages found containing the keyword '%s'\n", keyword)
			} else {
				fmt.Println("Search results:")
				for _, result := range results {
					fmt.Printf("[%s] %s: %s\n", result.Time, result.User, result.Content)
				}
			}
		} else {
			sendMessage(user, input)
		}

		fmt.Println("------------------")
	}
}