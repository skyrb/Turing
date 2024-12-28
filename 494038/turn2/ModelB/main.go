package main

import (
	"fmt"
	"text/template"
	"bytes"
)

// User represents the details of a user
type User struct {
	Name    string
	Email   string
	Gift    string
}

// generateThankYouEmail generates a personalized thank-you email using a template
func generateThankYouEmail(user User, templateName string) string {
	// Load the template from a file
	t, err := template.ParseFiles(fmt.Sprintf("templates/%s.txt", templateName))
	if err != nil {
		panic(err)
	}

	var emailBuffer bytes.Buffer
	// Execute the template with user data
	err = t.Execute(&emailBuffer, user)
	if err != nil {
		panic(err)
	}

	return emailBuffer.String()
}

func main() {
	// Example users
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Gift: "a book"},
		{Name: "Bob", Email: "bob@example.com", Gift: "a mug"},
		{Name: "Charlie", Email: "charlie@example.com", Gift: "a pen"},
	}

	// Generate and print thank-you emails for each user using the "thankyou" template
	for _, user := range users {
		email := generateThankYouEmail(user, "thankyou")
		fmt.Printf("Email to %s:\n%s\n\n", user.Email, email)
	}
} 