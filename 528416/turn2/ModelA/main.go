package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// filterEven logs and filters even numbers from a slice of integers.
func filterEven(numbers []int) []int {
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	log.Printf("Filtered even numbers: %v", evens)
	return evens
}

// reverseSlice logs and reverses a slice of integers.
func reverseSlice(numbers []int) []int {
	reversed := make([]int, len(numbers))
	copy(reversed, numbers)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	log.Printf("Reversed slice: %v", reversed)
	return reversed
}

// parseInput converts a slice of strings to a slice of integers.
func parseInput(args []string) ([]int, error) {
	var numbers []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", arg)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// main is the entry point of the program.
func main() {
	// Set up logging to output to the console.
	log.SetOutput(os.Stdout)

	// Validate command-line arguments.
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <operation> <list of numbers>")
		fmt.Println("Example: go run main.go sum 5 10 15 20")
		fmt.Println("Supported operations: sum, average, filterEven, reverse")
		os.Exit(1)
	}

	// Get the operation and input numbers.
	operation := os.Args[1]
	numberArgs := os.Args[2:]

	// Parse the input numbers.
	numbers, err := parseInput(numberArgs)
	if err != nil {
		log.Fatalf("Error parsing numbers: %v", err)
	}

	// Perform the requested operation.
	switch strings.ToLower(operation) {
	case "sum":
		sum := calculateSum(numbers)
		fmt.Printf("Sum: %d\n", sum)
	case "average":
		average := calculateAverage(numbers)
		fmt.Printf("Average: %.2f\n", average)
	case "filtereven":
		evens := filterEven(numbers)
		fmt.Printf("Filtered Even Numbers: %v\n", evens)
	case "reverse":
		reversed := reverseSlice(numbers)
		fmt.Printf("Reversed Slice: %v\n", reversed)
	default:
		log.Fatalf("Invalid operation: %s. Supported operations: sum, average, filterEven, reverse", operation)
	}
}