
package main  
import (  
    "fmt"
)  

type Case struct {  
    CaseID    string
    CaseName  string  
}  

type UserData struct {  
    Cases map[string]Case
}  

var justiceData = make(map[string]UserData)  

func main() {  
    addCase("user1", "Case 1.1", "Criminal case")  
    addCase("user1", "Case 1.2", "Civil case")  
    addCase("user2", "Case 2.1", "Family case")  

    printCases("user1")  
    printCases("user2")  
}  

func addCase(userID, caseID, caseName string) {  
    // Check if the user already exists in the justiceData map  
    if _, ok := justiceData[userID]; !ok {  
        // If the user doesn't exist, initialize a new UserData struct  
        justiceData[userID] = UserData{  
            Cases: make(map[string]Case),  
        }  
    }  
    // Add the new case to the user's case map  
    justiceData[userID].Cases[caseID] = Case{  
        CaseID:    caseID,  
        CaseName:  caseName,  
    }  
}  

func printCases(userID string) {  
    if userData, ok := justiceData[userID]; ok {  
        fmt.Println("Cases for User:", userID)  
        for caseID, caseData := range userData.Cases {  
            fmt.Printf("\tCase ID: %s, Case Name: %s\n", caseID, caseData.CaseName)  
        }  
    } else {  
        fmt.Println("User", userID, "not found.")  
    }  
}  
