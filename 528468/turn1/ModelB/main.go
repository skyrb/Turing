package main

import (
	"database/sql"
	"fmt"
)

func getUserByUsername(db *sql.DB, username string) (string, error) {
	query := fmt.Sprintf("SELECT name FROM users WHERE username = '%s'", sql.EscapeString(username))

	var name string
	err := db.QueryRow(query).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}