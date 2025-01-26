package main

import (
	"bufio"
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

// parseInput converts a string of space-separated numbers to a slice of integers.
func parseInput(input string) ([]int, error) {
	var numbers []int
	parts := strings.Fields(input)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// readInput reads user input from the console.
func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// main is the entry point of the program.
func main() {
	// Set up logging to output to the console.
	log.SetOutput(os.Stdout)

	for {
		// Prompt user for operation.
		operation := readInput("Choose an operation (sum, average, filterEven, reverse, quit): ")

		// Check if the user wants to quit.
		if strings.ToLower(operation) == "quit" {
			fmt.Println("Exiting the program.")
			break
		}

		// Prompt user for input numbers.
		input := readInput("Enter a list of numbers separated by spaces: ")

		// Parse the input numbers.
		numbers, err := parseInput(input)
		if err != nil {
			log.Printf("Error parsing numbers: %v", err)
			continue
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
			fmt.Println("Invalid operation. Please try again.")
		}
	}
}