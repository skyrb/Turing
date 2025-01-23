package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	// Maps command names to their corresponding functions
	commandMap = map[string]func(){
		"add":    addTask,
		"delete": deleteTask,
		"list":   listTasks,
		"exit":   exitApp,
	}

	// Suggestions for commonly misspelled commands
	commandSuggestions = map[string]string{
		"ad":   "add",
		"del":  "delete",
		"lis":  "list",
		"exi":  "exit",
		"q":    "exit",
		"quit": "exit",
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command: ")
		if !scanner.Scan() {
			break // Handle Ctrl+D or EOF
		}
		command := strings.TrimSpace(scanner.Text())
		executeCommand(command)
	}
}

func executeCommand(command string) {
	taskFunc, ok := commandMap[command]
	if ok {
		taskFunc() // Execute the corresponding task function
	} else {
		// Handle invalid command
		handleInvalidCommand(command)
	}
}

func handleInvalidCommand(command string) {
	fmt.Printf("Invalid command: '%s'\n", command)

	// Suggest valid options if the command has a suggestion
	suggestion, hasSuggestion := commandSuggestions[command]
	if hasSuggestion {
		fmt.Printf("Did you mean '%s'?\n", suggestion)
	}

	// Display available commands
	fmt.Println("Available commands:")
	for key := range commandMap {
		fmt.Println("  ", key)
	}
}

func addTask() {
	fmt.Println("Adding a new task...")
}

func deleteTask() {
	fmt.Println("Deleting a task...")
}

func listTasks() {
	fmt.Println("Listing tasks...")
}

func exitApp() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}