package main

import (
	"fmt"
	"strings"
)

func generateDynamicSQLQuery(username, email, role string) string {
	var conditions []string

	if username != "" {
		escapedUsername := strings.ReplaceAll(username, "'", "''")
		conditions = append(conditions, fmt.Sprintf("username = '%s'", escapedUsername))
	}

	if email != "" {
		escapedEmail := strings.ReplaceAll(email, "'", "''")
		conditions = append(conditions, fmt.Sprintf("email = '%s'", escapedEmail))
	}

	if role != "" {
		escapedRole := strings.ReplaceAll(role, "'", "''")
		conditions = append(conditions, fmt.Sprintf("role = '%s'", escapedRole))
	}

	var query string
	if len(conditions) > 0 {
		query = fmt.Sprintf("SELECT * FROM users WHERE %s;", strings.Join(conditions, " AND "))
	} else {
		query = "SELECT * FROM users;"
	}

	return query
}

func main() {
	query1 := generateDynamicSQLQuery("", "user@example.com", "admin")
	fmt.Println("Query 1:")
	fmt.Println(query1)

	query2 := generateDynamicSQLQuery("O'Reilly", "", "")
	fmt.Println("\nQuery 2:")
	fmt.Println(query2)

	query3 := generateDynamicSQLQuery("", "", "")
	fmt.Println("\nQuery 3:")
	fmt.Println(query3)
}