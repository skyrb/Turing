package main

import (
	"fmt"
)

// Case struct to hold details of a single case
type Case struct {
	CaseID    string
	CaseName  string
	Description string
}

// JusticeSystem struct encapsulates user data
type JusticeSystem struct {
	users map[string]UserData
}

// UserData struct encapsulates data for a single user
type UserData struct {
	cases map[string]Case
}

var justiceSystem = JusticeSystem{users: make(map[string]UserData)}

func (js *JusticeSystem) AddCase(userID, caseID, caseName, description string) {
	if _, ok := js.users[userID]; !ok {
		js.users[userID] = UserData{cases: make(map[string]Case)}
	}
	js.users[userID].cases[caseID] = Case{
		CaseID:    caseID,
		CaseName:  caseName,
		Description: description,
	}
}

func (js *JusticeSystem) GetUserCases(userID string) ([]Case, error) {
	if userData, ok := js.users[userID]; ok {
		var cases []Case
		for _, caseData := range userData.cases {
			cases = append(cases, caseData)
		}
		return cases, nil
	} else {
		return nil, fmt.Errorf("user not found: %s", userID)
	}
}

func main() {
	justiceSystem.AddCase("user1", "Case 1.1", "Criminal case", "Details of Case 1.1 for user1")
	justiceSystem.AddCase("user1", "Case 1.2", "Civil case", "Details of Case 1.2 for user1")
	justiceSystem.AddCase("user2", "Case 2.1", "Family case", "Details of Case 2.1 for user2")

	user1Cases, err := justiceSystem.GetUserCases("user1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cases for User: user1")
		for _, caseData := range user1Cases {
			fmt.Printf("\tCase ID: %s, Case Name: %s, Description: %s\n", caseData.CaseID, caseData.CaseName, caseData.Description)
		}
	}

	user2Cases, err := justiceSystem.GetUserCases("user2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cases for User: user2")
		for _, caseData := range user2Cases {
			fmt.Printf("\tCase ID: %s, Case Name: %s, Description: %s\n", caseData.CaseID, caseData.CaseName, caseData.Description)
		}
	}

	// Attempting to get cases for a non-existing user will return an error
	nonExistentUserCases, err := justiceSystem.GetUserCases("user3")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cases for User: user3")
		for _, caseData := range nonExistentUserCases {
			fmt.Printf("\tCase ID: %s, Case Name: %s, Description: %s\n", caseData.CaseID, caseData.CaseName, caseData.Description)
		}
	}
}