package main  
import (  
    "fmt"
    "time"
)

type UserState struct {  
    Timestamp time.Time
    Users     []string
    Changes   []string
}

func main() {  
    userStates := make([]UserState, 0)

    currentUsers := []string{"User1", "User2", "User3"}
    userStates = append(userStates, UserState{
        Timestamp: time.Now(),
        Users:     currentUsers,
        Changes:   []string{"Initial state"},
    })

    fmt.Println("Initial state:", currentUsers)

    // User logs in
    currentUsers = append(currentUsers, "User4")
    userStates = append(userStates, UserState{
        Timestamp: time.Now(),
        Users:     currentUsers,
        Changes:   []string{"User4 logged in"},
    })

    fmt.Println("After User4 logs in:", currentUsers)

    // User logs out
    for i, user := range currentUsers {
        if user == "User3" {
            currentUsers = append(currentUsers[:i], currentUsers[i+1:]...)
            userStates = append(userStates, UserState{
                Timestamp: time.Now(),
                Users:     currentUsers,
                Changes:   []string{"User3 logged out"},
            })
            break
        }
    }

    fmt.Println("After User3 logs out:", currentUsers)

    // Display all states
    fmt.Println("\nUser States:")
    for i, state := range userStates {
        fmt.Printf("State %d: %v, Users: %v, Changes: %v\n", i+1, state.Timestamp, state.Users, state.Changes)
    }
}