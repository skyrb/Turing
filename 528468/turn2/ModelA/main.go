package main

import (
	"fmt"
	"strings"
)

// escapeSingleQuotes escapes single quotes in the input string to prevent syntax errors.
func escapeSingleQuotes(input string) string {
	return strings.ReplaceAll(input, "'", "''")
}

// generateQuery creates a SQL query based on the provided conditions.
func generateQuery(table string, conditions map[string]string) string {
	// Start constructing the query
	baseQuery := fmt.Sprintf("SELECT * FROM %s WHERE 1=1", table)

	// Add conditions to the query
	for key, value := range conditions {
		// Escape the value to handle single quotes
		escapedValue := escapeSingleQuotes(value)
		condition := fmt.Sprintf(" AND %s = '%s'", key, escapedValue)
		baseQuery += condition
	}

	return baseQuery
}

func main() {
	// Define conditions
	conditions := map[string]string{
		"username": "JohnDoe",
		"email":    "john.doe@example.com",
		"role":     "admin",
	}

	// Generate the query
	query := generateQuery("users", conditions)

	// Print the constructed query
	fmt.Println("Constructed Query:")
	fmt.Println(query)
}