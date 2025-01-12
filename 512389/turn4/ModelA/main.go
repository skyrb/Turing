package main

import (
	"fmt"
	"sync"
)

type Case struct {
	CaseID      string
	CaseName    string
	Description string
}

var justiceSystem = &JusticeSystem{users: &sync.Map{}}

type JusticeSystem struct {
	users *sync.Map
}

func (js *JusticeSystem) AddCase(userID, caseID, caseName, description string) {
	js.users.Store(userID, func() interface{} {
		userData := UserData{cases: &sync.Map{}}
		userData.cases.Store(caseID, Case{
			CaseID:      caseID,
			CaseName:    caseName,
			Description: description,
		})
		return userData
	}())
}

func (js *JusticeSystem) GetUserCases(userID string) ([]Case, error) {
	var cases []Case

	if userDataInterface, ok := js.users.Load(userID); ok {
		userData := userDataInterface.(UserData)
		userData.cases.Range(func(caseID, caseData interface{}) bool {
			caseObj := caseData.(Case)
			cases = append(cases, caseObj)
			return true // Continue the iteration
		})
	} else {
		return nil, fmt.Errorf("user not found: %s", userID)
	}

	return cases, nil
}

func main() {
	// Launch a few goroutines to simulate concurrent user operations
	go addAndGetCases("user1")
	go addAndGetCases("user2")
	go addAndGetCases("user1")

	// Wait for goroutines to complete
	var wg sync.WaitGroup
	wg.Add(3)
	go func() { addAndGetCases("user1"); wg.Done() }()
	go func() { addAndGetCases("user2"); wg.Done() }()
	go func() { addAndGetCases("user3"); wg.Done() }()
	wg.Wait()
}

func addAndGetCases(userID string) {
	justiceSystem.AddCase(userID, "Case 1.1", "Criminal case", "Details of Case 1.1 for user1")
	justiceSystem.AddCase(userID, "Case 1.2", "Civil case", "Details of Case 1.2 for user1")

	userCases, err := justiceSystem.GetUserCases(userID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cases for User:", userID)
		for _, caseData := range userCases {
			fmt.Printf("\tCase ID: %s, Case Name: %s, Description: %s\n", caseData.CaseID, caseData.CaseName, caseData.Description)
		}
	}
}