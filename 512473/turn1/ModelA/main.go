package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// Create a slice of users
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// Format and print the users slice
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
	}

	// Alternatively, use a formatted string literal
	for _, user := range users {
		fmt.Println(fmt.Sprintf("Name: %s, Age: %d", user.Name, user.Age))
	}
}