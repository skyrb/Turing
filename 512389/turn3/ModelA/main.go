package main

import (
	"fmt"
	"sync"
)

type Case struct {
	CaseID    string
	CaseName  string
	Description string
}

type JusticeSystem struct {
	users  map[string]UserData
	mu     sync.RWMutex // Main lock for managing user entries
	locks  map[string]*sync.Mutex // Lock per user
}

type UserData struct {
	cases map[string]Case
	mu    *sync.Mutex
}

var justiceSystem = &JusticeSystem{users: make(map[string]UserData), locks: make(map[string]*sync.Mutex)}

func (js *JusticeSystem) EnsureLock(userID string) {
	js.mu.Lock()
	defer js.mu.Unlock()
	if js.locks[userID] == nil {
		js.locks[userID] = new(sync.Mutex)
	}
}

func (js *JusticeSystem) AddCase(userID, caseID, caseName, description string) {
	js.EnsureLock(userID)
	js.locks[userID].Lock()
	defer js.locks[userID].Unlock()
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
	js.EnsureLock(userID)
	js.locks[userID].RLock()
	defer js.locks[userID].RUnlock()

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
	// Launch a few goroutines to simulate concurrent user operations
	go addAndGetCases("user1")
	go addAndGetCases("user2")
	go addAndGetCases("user1")
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