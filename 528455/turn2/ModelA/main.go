package main

import (
	"fmt"
)

// Person struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// String method to satisfy the fmt.Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// formatMessage constructs a unified message from a variety of input types
func formatMessage(args ...interface{}) string {
	var message string
	for _, arg := range args {
		message += fmt.Sprintf("%v, ", arg)
	}
	// Removing the last comma and space for cleaner output
	if len(message) > 0 {
		message = message[:len(message)-2]
	}
	return message
}

func main() {
	// Create an instance of Person
	person := Person{Name: "John Doe", Age: 30}

	// Construct a message using different data types
	result := formatMessage("Welcome", 123, person, 45.67, true)

	// Print the result
	fmt.Println(result)
}