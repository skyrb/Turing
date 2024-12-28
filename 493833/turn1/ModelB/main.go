package main

import (
	"fmt"
	"strings"
)

func cleanData(input string) string {
	// Remove leading/trailing whitespaces
	cleaned := strings.TrimSpace(input)

	// Replace multiple spaces with a single space
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")

	return cleaned
}

func generateReport(data string) string {
	header := fmt.Sprintf("Report for '%s'\n", data)
	body := fmt.Sprintf("Data: %s\n", data)
	return header + body
}

func main() {
	rawData := " Hello, World!  \n This is  a   test. "

	// Chaining of functions to clean and generate the report
	report := generateReport(cleanData(rawData))

	fmt.Println(report)
}