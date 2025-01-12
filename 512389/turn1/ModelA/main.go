package main

import (
	"fmt"
)

// Define a struct to represent a case
type Case struct {
	CaseID string `json:"case_id"`
	Title   string `json:"title"`
	Description string `json:"description"`
}

// Define a struct to represent a user with their cases
type User struct {
	Cases map[string]Case `json:"cases"`
}

func main() {
	// Create a map to hold users
	users := make(map[string]User)

	// Add two users to the map
	users["user1"] = User{
		Cases: map[string]Case{
			"case1": Case{CaseID: "case1", Title: "Case One for User One", Description: "Description of case one."},
			"case2": Case{CaseID: "case2", Title: "Case Two for User One", Description: "Description of case two."},
		},
	}

	users["user2"] = User{
		Cases: map[string]Case{
			"case3": Case{CaseID: "case3", Title: "Case One for User Two", Description: "Description of case one."},
		},
	}

	// Display cases for user1
	fmt.Println("Cases for user1:")
	for caseID, caseData := range users["user1"].Cases {
		fmt.Printf("CaseID: %s, Title: %s, Description: %s\n", caseID, caseData.Title, caseData.Description)
	}

	// Display cases for user2
	fmt.Println("\nCases for user2:")
	for caseID, caseData := range users["user2"].Cases {
		fmt.Printf("CaseID: %s, Title: %s, Description: %s\n", caseID, caseData.Title, caseData.Description)
	}
}