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

// JusticeSystem represents the entire justice system data
type JusticeSystem struct {
	users sync.Map // Map of userID -> UserData
}

// UserData struct holds data for a single user
type UserData struct {
	cases sync.Map // Map of caseID -> Case
}

var justiceSystem = &JusticeSystem{}

func (js *JusticeSystem) AddCase(userID, caseID, caseName, description string) {
	userData, _ := js.users.LoadOrStore(userID, UserData{})
	ud := userData.(UserData)
	ud.cases.Store(caseID, Case{
		CaseID:      caseID,
		CaseName:    caseName,
		Description: description,
	})
	js.users.Store(userID, ud)
}

func (js *JusticeSystem) GetUserCases(userID string) ([]Case, error) {
	userData, ok := js.users.Load(userID)
	if !ok {
		return nil, fmt.Errorf("user not found: %s", userID)
	}
	ud := userData.(UserData)
	var cases []Case
	ud.cases.Range(func(key, value interface{}) bool {
		cases = append(cases, value.(Case))
		return true
	})
	return cases, nil
}

func main() {
	// Launch multiple goroutines for concurrent operations
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			addAndGetCases(fmt.Sprintf("user%d", i))
		}(i)
	}
	wg.Wait()
}

func addAndGetCases(userID string) {
	justiceSystem.AddCase(userID, "Case1", "Criminal case", "Details of Case1 for "+userID)
	justiceSystem.AddCase(userID, "Case2", "Civil case", "Details of Case2 for "+userID)

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