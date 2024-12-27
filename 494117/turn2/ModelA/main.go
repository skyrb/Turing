package main

import (
	"fmt"
)

// A simple counter structure
type Counter struct {
	value int
}

func (c *Counter) Increment() error {
	// Back up the original value
	originalValue := c.value
	
	// Perform some work that might fail
	fmt.Println("Incrementing counter...")
	if c.value == 5 {
		return fmt.Errorf("counter exceeds maximum value of 5")
	}

	// Increment the counter
	c.value++

	defer func() {
		// Rollback in case of error
		if err := recover(); err != nil {
			fmt.Printf("Rolling back change due to error: %v\n", err)
			c.value = originalValue
		}
	}()

	// Simulate a "successful" increment
	return nil
}

func main() {
	counter := Counter{value: 4}

	if err := counter.Increment(); err != nil {
		fmt.Println("Increment failed:", err)
	} else {
		fmt.Println("Increment successful.")
	}

	fmt.Println("Final counter value:", counter.value)
}