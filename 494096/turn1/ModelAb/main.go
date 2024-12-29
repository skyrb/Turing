package main

import (
	"fmt"
	"time"
)

func calculateSum(a, b int) int {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("calculateSum(%d, %d) took %s\n", a, b, end.Sub(start))
	}()
	return a + b
}

func main() {
	fmt.Println(calculateSum(1000000, 2000000))
}