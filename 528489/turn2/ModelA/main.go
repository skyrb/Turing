package main

import (
	"fmt"
	"os"

	"github.com/agnivade/levenshtein"
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [arguments]")
		fmt.Println("Available commands: add, delete, list")
		os.Exit(1)
	}

	command := os.Args[1]

	if taskFunc, ok := commandMap[command]; ok {
		taskFunc()
	} else {
		fmt.Printf("Invalid command: '%s'\n", command)
		suggestion := suggestCommand(command)
		if suggestion != "" {
			fmt.Printf("Did you mean '%s'?\n", suggestion)
		}
		os.Exit(1)
	}
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

func suggestCommand(input string) string {
	minDistance := len(input) + 1
	bestMatch := ""

	for cmd := range commandMap {
		distance := levenshtein.ComputeDistance(input, cmd)
		if distance < minDistance {
			minDistance = distance
			bestMatch = cmd
		}
	}

	if minDistance <= 2 { // Example threshold for suggestion
		return bestMatch
	}

	return ""
}