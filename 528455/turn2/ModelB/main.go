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

// FormatMessage function takes various data types as arguments and formats them into a unified message using fmt.Sprintf
func FormatMessage(parts ...interface{}) string {
	var msg string
	for _, part := range parts {
		if str, ok := part.(string); ok {
			// If it's a string, append it directly
			msg += str
		} else if num, ok := part.(int); ok {
			// If it's an int, append it as a string
			msg += fmt.Sprintf("%d", num)
		} else if person, ok := part.(Person); ok {
			// If it's a Person struct, call its String() method
			msg += person.String()
		} else {
			// For other types, just print them as a string representation
			msg += fmt.Sprintf("%v", part)
		}
		// Add a space after each part
		msg += " "
	}
	return msg
}

func main() {
	// Test the FormatMessage function
	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{Name: "Bob", Age: 30}

	message := FormatMessage(
		"Hello, ",
		"how are you?",
		42,
		person1,
		"today is",
		"Thursday",
		person2,
		"?",
	)

	fmt.Println(message)
}  