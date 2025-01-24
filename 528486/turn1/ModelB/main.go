package main  
import (  
    "fmt"
)  
func main() {  
    // Civic engagement data  
    data := map[string][]string{  
        "User1": {"Option A", "Option B"},  
        "User2": {"Option B", "Option C"},  
        "User3": {"Option A"},
        "User4": {"Option C"},
        "User5": {"Option A", "Option B", "Option C"},  
    }  
    voteCount := make(map[string]int)  
    // Iterate through the data and count the votes for each option  
    for _, options := range data {  
        for _, option := range options {  
            voteCount[option]++  
        }  
    }  
    // Display the vote count for each option  
    fmt.Println("Vote Count:")  
    for option, count := range voteCount {  
        fmt.Printf("%s: %d\n", option, count)  
    }  
}  