package main

import (
	"fmt"
	"time"
)

// parseAndProcess attempts to parse a string and return the processed result.
func parseAndProcess(input string) (int, error) {
	// Simulate some expensive parsing logic
	time.Sleep(time.Millisecond * 100)
	
	num, err := fmt.Sscanf(input, "%d", &num)
	if err != nil {
		return 0, err
	}
	
	if num == 0 {
		return 0, fmt.Errorf("number is zero")
	}
	
	return num, nil
}

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		_, _ = parseAndProcess(fmt.Sprintf("%d", i))
	}
	end := time.Now()
	
	fmt.Println("Time taken:", end.Sub(start))
}