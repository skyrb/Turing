package main

import (
	"bytes"
	"fmt"
	"strings"
)

func formatStringsEfficiently(data []string) []string {
	var results []string
	var buffer bytes.Buffer
	maxLength := 0

	// Determine the maximum string length
	for _, name := range data {
		if len(name) > maxLength {
			maxLength = len(name)
		}
	}

	// Create a static prefix or suffix to use once
	staticPrefix := "Name: "
	staticSuffix := "\n"

	// Format strings in batches
	for _, name := range data {
		// Calculate padding required for the string
		padding := maxLength - len(name)
		paddingStr := strings.Repeat(" ", padding)

		// Write formatted string to the buffer
		buffer.WriteString(staticPrefix)
		buffer.WriteString(name)
		buffer.WriteString(paddingStr)
		buffer.WriteString(staticSuffix)
	}

	// Append all formatted strings to the results slice
	results = append(results, buffer.String())

	return results
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David"}
	formatted := formatStringsEfficiently(names)
	fmt.Println(formatted[0])
}