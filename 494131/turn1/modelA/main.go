package main

import (
	"fmt"
)

// sumAndError returns the sum of two integers and an error if the sum exceeds 100.
func sumAndError(a, b int) (int, error) {
	result := a + b
	if result > 100 {
		return 0, fmt.Errorf("sum is greater than 100: %d + %d = %d", a, b, result)
	}
	return result, nil
}

func main() {
	// Example 1: Successful sum
	result, err := sumAndError(50, 40)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", result)
	}

	// Example 2: Sum exceeds 100
	result, err = sumAndError(60, 50)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", result)
	}
}