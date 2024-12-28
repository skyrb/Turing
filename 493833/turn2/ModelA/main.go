package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Sample data
	data := " 123 Main St, Suite 456, Anytown, USA 12345 "

	// Parse and reformat the address using fmt.Sprintf with additional methods
	cleanedAddress := fmt.Sprintf(
		"%s\n%s", // Use cleaned address and extracted zip code
		strings.TrimSpace(
			strings.Join(
				// Split by commas and reformat
				strings.Map(func(r rune) rune {
					if r == ',' {
						return '\n'
					}
					return r
				}, strings.Split(strings.TrimSpace(data), ",")),
				" ",
			),
		),
		extractZipCode(data),
	)

	// Generate a formatted report
	report := fmt.Sprintf(
		"Address:\n%s\n",
		cleanedAddress,
	)

	fmt.Println(report)
}

// Function to extract zip code using a regular expression
func extractZipCode(input string) string {
	zipCodeRegex := regexp.MustCompile(`(\d{5})`) // Matches a 5-digit zip code
	match := zipCodeRegex.FindString(input)
	return match
}