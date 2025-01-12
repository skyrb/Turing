package main  
import (  
    "fmt"
    "sync"
)  

type Case struct {  
    CaseID    string  
    CaseName  string  
}  

type JusticeApp struct {  
    userData map[string]map[string]Case  
    mu sync.Mutex  
}  

func NewJusticeApp() *JusticeApp {  
    return &JusticeApp{  
        userData: make(map[string]map[string]Case),  
    }  
}  

func (ja *JusticeApp) AddCase(userID, caseID, caseName string) {  
    ja.mu.Lock()  
    defer ja.mu.Unlock()  

    // Initialize user map if it doesn't exist  
    if _, ok := ja.userData[userID]; !ok {  
        ja.userData[userID] = make(map[string]Case)  
    }  

    // Add the case  
    ja.userData[userID][caseID] = Case{CaseID: caseID, CaseName: caseName}  
}  

func (ja *JusticeApp) GetCase(userID, caseID string) (Case, bool) {  
    ja.mu.Lock()  
    defer ja.mu.Unlock()  

    // Check if user and case exist  
    userCases, ok := ja.userData[userID]  
    if !ok {  
        return Case{}, false // User not found  
    }  

    caseData, found := userCases[caseID]  
    return caseData, found  
}  

func (ja *JusticeApp) PrintAllCasesForUser(userID string) {  
    ja.mu.Lock()  
    defer ja.mu.Unlock()  

    fmt.Println("Cases for User:", userID)  
    userCases, ok := ja.userData[userID]  
    if !ok {  
        fmt.Println("User", userID, "not found.")  
        return  
    }  

    for caseID, caseData := range userCases {  
        fmt.Printf("\tCase ID: %s, Case Name: %s\n", caseID, caseData.CaseName)  
    }  
}  


func main() {  
    justiceApp := NewJusticeApp()  

    justiceApp.AddCase("user1", "Case 1.1", "Criminal case")  
    justiceApp.AddCase("user1", "Case 1.2", "Civil case")  
    justiceApp.AddCase("user2", "Case 2.1", "Family case")  

    justiceApp.PrintAllCasesForUser("user1")  
    justiceApp.PrintAllCasesForUser("user2")  

    // Attempting to access user2's case via user1's namespace will fail
    _, found := justiceApp.GetCase("user1", "Case 2.1")  
    fmt.Println("Case 2.1 for user1 found:", found) // Output: Case 2.1 for user1 found: false
}  

