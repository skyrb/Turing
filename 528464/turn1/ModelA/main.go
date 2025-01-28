package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Replace with your database driver
)

func main() {
	age := 30
	location := "New York"

	// Open connection to the database (example using PostgreSQL)
	db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Construct the SQL query using fmt.Sprintf
	query := fmt.Sprintf("SELECT * FROM users WHERE age > %d AND location = '%s'", age, location)

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Example: Iterate over the rows and process the results
	for rows.Next() {
		// Assuming users table has columns id, name, age, location
		var id int
		var name string
		var age int
		var location string

		err := rows.Scan(&id, &name, &age, &location)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s, Age: %d, Location: %s\n", id, name, age, location)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}