package main

import (
	"fmt"
)

func main() {
	// Define the data structure
	type Task struct {
		Name    string
		Status string
	}

	project := "Smart Home System"
	owner := "Alice Johnson"
	tasks := []Task{
		{"Design", "Completed"},
		{"Development", "In Progress"},
		{"Testing", "Pending"},
	}

	// Create the formatted string
	output := fmt.Sprintf("Project: %s\nOwner: %s\nTasks:\n", project, owner)
	for _, task := range tasks {
		output += fmt.Sprintf("    - %s: %s\n", task.Name, task.Status)
	}

	fmt.Print(output)
}