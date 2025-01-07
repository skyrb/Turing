package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Your application code here

	// Periodically check the number of goroutines
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
		}
	}
}