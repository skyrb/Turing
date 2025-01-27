package main

import (
	"fmt"
)

// Person struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// String method to implement the fmt.Stringer interface
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	// Creating an instance of Person
	person := Person{
		Name: "Alice",
		Age:  30,
	}

	// Using fmt.Sprintf to get a formatted string
	result := fmt.Sprintf("%v", person)
	fmt.Println(result) // Output: Name: Alice, Age: 30

	// Alternatively, using fmt.Println directly
	fmt.Println(person) // Output: Name: Alice, Age: 30
}