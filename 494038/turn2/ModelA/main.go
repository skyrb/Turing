package main

import (
	"fmt"
	"text/template"
	"os"
)

// User represents the details of a user
type User struct {
	Name    string
	Email   string
	Gift    string
}

// getTemplateCompiled compiles the email template from a file
func getTemplateCompiled() *template.Template {
	// Compile the template from a string directly
	emailTemplate := `
Dear {{.Name}},

Thank you so much for your thoughtful gift of {{.Gift}}. I really appreciate it!

Looking forward to our next meeting.

Warm regards,
[Your Name]
`

	// Creating a new template and parsing the text
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		fmt.Printf("Error compiling template: %v\n", err)
		return nil
	}
	return tmpl
}

func main() {
	// Example users
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Gift: "a book"},
		{Name: "Bob", Email: "bob@example.com", Gift: "a mug"},
		{Name: "Charlie", Email: "charlie@example.com", Gift: "a pen"},
	}

	// Get the compiled template
	tmpl := getTemplateCompiled()
	if tmpl == nil {
		fmt.Println("Failed to compile template, exiting.")
		return
	}

	// Generate and print thank-you emails for each user
	for _, user := range users {
		email := &user
		err := tmpl.Execute(os.Stdout, email)
		if err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			continue
		}
		fmt.Println("\nEmail to", user.Email, ":\n---------------------")
	}
}