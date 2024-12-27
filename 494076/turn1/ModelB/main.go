package main

import (
	"fmt"
)

// MyError is a custom error type representing invalid input.
type MyError struct {
	Reason string
	Value  int
}

func (e *MyError) Error() string {
	// Optionally, you can call fmt.Sprintf here to include detailed fields
	// in the error message.
	return fmt.Sprintf("invalid input: %s (value: %d)", e.Reason, e.Value)
}

func main() {
	// Demonstrate custom error formatting
	err := &MyError{Reason: "value is too large", Value: 100}
	fmt.Println(err) // Output: invalid input: value is too large (value: 100)
}