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

// JusticeSystem struct encapsulates user data
type JusticeSystem struct {
	users map[string]UserData
}

// UserData struct encapsulates data for a single user
type UserData struct {
	cases map[string]Case
}

var justiceSystem = NewJusticeSystem()

func NewJusticeSystem() *JusticeSystem {  
    js := &JusticeSystem{  
        users: make(map[string]UserData),
        locks: make(map[string]*sync.RWMutex),  
    }  
    return js  
}  

func (js *JusticeSystem) GetUserLock(userID string) *sync.RWMutex {  
    js.mu.RLock()  
    defer js.mu.RUnlock()  
    if lock, ok := js.locks[userID]; ok {  
        return lock  
    }  
    js.mu.RUnlock()  
    js.mu.Lock()  
    defer js.mu.Unlock()  
    if lock, ok := js.locks[userID]; ok {  
        return lock  
    }  
    lock := &sync.RWMutex{}  
    js.locks[userID] = lock  
    return lock  
}  

func (js *JusticeSystem) AddCase(userID, caseID, caseName, description string) {  
    lock := js.GetUserLock(userID)  
    lock.Lock()  
    defer lock.Unlock()  
    // Rest of the addCase function remains the same...
}  

func (js *JusticeSystem) GetUserCases(userID string) ([]Case, error) {  
    lock := js.GetUserLock(userID)  
    lock.RLock()  
    defer lock.RUnlock()  
    // Rest of the GetUserCases function remains the same...
}  

func main() {  
    go func() {  
        justiceSystem.AddCase("user1", "Case 1.1", "Criminal case", "Details of Case 1.1 for user1")  
    }()  
  
    go func() {  
        justiceSystem.AddCase("user2", "Case 2.1", "Family case", "Details of Case 2.1 for user2")  
    }()  

    go func() {  
        cases, err := justiceSystem.GetUserCases("user1")  
        if err != nil {  
            fmt.Println("Error:", err)  
        } else {  
            fmt.Println("Cases for User: user1")  
            for _, caseData := range cases {  
                fmt.Printf("\tCase ID: %s, Case Name: %s, Description: %s\n", caseData.CaseID, caseData.CaseName, caseData.Description)  
            }  
        }  
    }()  
    // ... more goroutines adding cases and accessing user data
} 