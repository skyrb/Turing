package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Sample data with different address formats
	data1 := "123 Main St, Suite 456, Anytown, USA 12345"
	data2 := "456 Elm Ave apt 7B, Cityville, XY 78901-2345"
	data3 := "Another Road, Part VIII, City, ZA 1112 3456"

    // Base format to clean the address
    addressBaseFormat := "%s %s, %s, %s %s"

	cleanAddresses := map[string]string{
		data1: processAddress(data1, addressBaseFormat),
		data2: processAddress(data2, addressBaseFormat),
		data3: processAddress(data3, addressBaseFormat),
	}

	// Display the cleaned addresses
	for data, cleaned := range cleanAddresses {
		fmt.Printf("Input Address: %s\nCleaned Address: %s\n\n", data, cleaned)
	}
}

func processAddress(address, baseFormat string) string {
	// Split the address by commas to separate street, town, state, zip
	parts := strings.Split(address, ",")

	street := strings.TrimSpace(parts[0])
	city := strings.TrimSpace(parts[1])
	state := extractState(strings.TrimSpace(parts[2]))
	zip := extractZipCode(address)

	return fmt.Sprintf(baseFormat, street, state, city, zip)
}

func extractZipCode(address string) string {
	zipCodeRegex := regexp.MustCompile(`[0-9]{5}(?:-[0-9]{4})?`)
	match := zipCodeRegex.FindString(address)
	return strings.TrimSpace(match)
}

func extractState(statePart string) string {
	// Simple state abbreviation extraction using first 2 letters
	state := strings.TrimSpace(strings.ToUpper(statePart[:2]))
	// Regular expressions can be more comprehensive for more complex state formats
	return state
} 