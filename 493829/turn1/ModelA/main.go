package main

import (
	"errors"
	"fmt"
)

// User struct representing a user
type User struct {
	ID   int
	Name string
}

// FetchData interface for the data fetching function
type FetchData interface {
	Fetch() (interface{}, error)
}

// UserFetcher struct implementing the FetchData interface
type UserFetcher struct {
}

// Fetch method for the UserFetcher
func (uf *UserFetcher) Fetch() (interface{}, error) {
	// Simulate fetching data from the database
	queryResult := 1
	switch queryResult {
	case 1: // Single user
		return User{ID: 1, Name: "Alice"}, nil
	case 2: // Multiple users
		return []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}, nil
	case 3: // Error
		return nil, errors.New("Database error")
	default:
		return nil, errors.New("Unknown error")
	}
}

func main() {
	fetcher := &UserFetcher{}
	result, err := fetcher.Fetch()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch t := result.(type) {
	case User:
		fmt.Printf("Single user fetched: %+v\n", t)
	case []User:
		fmt.Printf("Multiple users fetched: %+v\n", t)
	default:
		fmt.Println("Unknown result type")
	}
}