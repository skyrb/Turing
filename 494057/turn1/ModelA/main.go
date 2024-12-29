package main

import (
	"fmt"
	"strings"
)

func main() {
	// Sample data
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	ages := []int{30, 22, 45, 35, 27}

	// Calculate the maximum length of the names for alignment
	maxNameLength := strings.Max(names...)

	// Format and print the header
	fmt.Printf("%*s | %*s\n", maxNameLength, "Name", 5, "Age")

	// Iterate over the data and print each entry
	for i, name := range names {
		// Calculate the width for the name based on the maximum length
		width := maxNameLength - len(name) + 2
		fmt.Printf("%*s | %d\n", width, name, ages[i])
	}
}