package main  

import (  
    "fmt"
    "os"
)

var (  
    // Maps command names to their corresponding functions
    commandMap = map[string]func(){  
        "add":    addTask,  
        "delete": deleteTask,  
        "list":   listTasks,  
    }  
)  

func main() {  
    // Check if an argument is provided after the executable name
    if len(os.Args) < 2 {  
        fmt.Println("Usage: go run main.go <command> [arguments]")  
        fmt.Println("Available commands: add, delete, list")  
        os.Exit(1)  
    }  

    // Extract the command from the first argument
    command := os.Args[1]  

    // Look up the corresponding function for the command
    taskFunc, ok := commandMap[command]  
    if !ok {  
        fmt.Println("Invalid command. Use 'go run main.go help' for available commands.")
        os.Exit(1)  
    }  

    // Call the task function to execute the respective action
    taskFunc()  
}  

func addTask() {  
    // Implement your add task functionality here
    fmt.Println("Adding a new task...")  
}  

func deleteTask() {  
    // Implement your delete task functionality here
    fmt.Println("Deleting a task...")  
}  

func listTasks() {  
    // Implement your list task functionality here
    fmt.Println("Listing tasks...")  
}  