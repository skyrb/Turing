package main

import (
    "database/sql"
    "fmt"
    "log"
	"strings"

    _ "github.com/go-sql-driver/mysql" // Import the driver
)

func main() {
    // Example username with a single quote
    username := "O'Reilly"

    // Escape single quotes by doubling them up
    safeUsername := fmt.Sprintf("'%s'", escapeSQL(username))

    // Construct the query
    query := fmt.Sprintf("SELECT * FROM users WHERE username = %s", safeUsername)

    // Assuming db is a previously opened database connection
    var db *sql.DB // This would normally be initialized properly
    rows, err := db.Query(query) 

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    // Handle the query results
    for rows.Next() {
        // Process each row
    }
}

// escapeSQL properly escapes single quotes
func escapeSQL(input string) string {
    // Replace single quote with two single quotes
    return strings.ReplaceAll(input, "'", "''")
}