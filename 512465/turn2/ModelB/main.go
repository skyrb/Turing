package main  
import (  
    "fmt"
    "strings"
)

type Message struct {  
    User    string
    Content string
}

type MessagingApp struct {  
    Messages []Message
}

func (app *MessagingApp) AddMessage(user, content string) {  
    app.Messages = append(app.Messages, Message{User: user, Content: content})
}

func (app *MessagingApp) SearchMessages(keywords ...string) []Message {  
    var results []Message
    keywordsLower := make([]string, len(keywords))

    // Convert all keywords to lowercase
    for i, keyword := range keywords {
        keywordsLower[i] = strings.ToLower(keyword)
    }

    for _, message := range app.Messages {  
        contentLower := strings.ToLower(message.Content)
        userLower := strings.ToLower(message.User)
        
        // Check if all keywords are present in either the User or Content field
        found := true
        for _, keyword := range keywordsLower {
            if !strings.Contains(contentLower, keyword) && !strings.Contains(userLower, keyword) {
                found = false
                break
            }
        }
        if found {
            results = append(results, message)
        }
    }
    return results
}  

func main() {  
    app := MessagingApp{}
    app.AddMessage("Alice", "Hi Bob! Let's meet at 3pm for lunch.")
    app.AddMessage("Bob", "Okay, Alice! See you then.")
    app.AddMessage("Charlie", "Just checking in. How are you?")
    app.AddMessage("David", "Hi Alice, can we have a meeting at 2pm?")
    app.AddMessage("Alice", "Sure, David! That sounds good.")

    // Search for messages containing both "Alice" and "meeting"
    keywords := []string{"Alice", "David"}
    results := app.SearchMessages(keywords...)

    fmt.Printf("Messages containing '%s': \n", strings.Join(keywords, " and "))
    for _, result := range results {  
        fmt.Printf("%s: %s\n", result.User, result.Content)
    }
}  