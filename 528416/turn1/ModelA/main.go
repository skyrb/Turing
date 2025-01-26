package main

import (
	"fmt"
	"log"
	"os"
)

// calculateSum logs and calculates the sum of the integers in a slice.
func calculateSum(numbers []int) int {
	log.Printf("Calculating sum: %v", numbers)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// calculateAverage logs and calculates the average of the integers in a slice.
func calculateAverage(numbers []int) float64 {
	log.Printf("Calculating average: %v", numbers)
	sum := calculateSum(numbers)
	if len(numbers) == 0 {
		return 0
	}
	return float64(sum) / float64(len(numbers))
}

// main is the entry point of the program.
func main() {
	// Set up logging to output to the console.
	log.SetOutput(os.Stdout)

	// Example list of numbers.
	numbers := []int{5, 10, 15, 20}

	// Calculate the sum of the numbers.
	sum := calculateSum(numbers)
	fmt.Printf("Sum: %d\n", sum)

	// Calculate the average of the numbers.
	average := calculateAverage(numbers)
	fmt.Printf("Average: %.2f\n", average)
}