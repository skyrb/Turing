package main

import (
	"fmt"
)

// User represents the details of a user
type User struct {
	Name    string
	Email   string
	Gift    string
}

// generateThankYouEmail generates a personalized thank-you email for a given user
func generateThankYouEmail(user User) string {
	// Email template with placeholders
	emailTemplate := `
Dear %s,

Thank you so much for your thoughtful gift of %s. I really appreciate it!

Looking forward to our next meeting.

Warm regards,
[Your Name]
`
	// Format the template with the user's details
	return fmt.Sprintf(emailTemplate, user.Name, user.Gift)
}

func main() {
	// Example users
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Gift: "a book"},
		{Name: "Bob", Email: "bob@example.com", Gift: "a mug"},
		{Name: "Charlie", Email: "charlie@example.com", Gift: "a pen"},
	}

	// Generate and print thank-you emails for each user
	for _, user := range users {
		email := generateThankYouEmail(user)
		fmt.Printf("Email to %s:\n%s\n\n", user.Email, email)
	}
}