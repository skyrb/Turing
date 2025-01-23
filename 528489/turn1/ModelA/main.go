package main

import (
	"fmt"
	"os"
	"strings"
)

var tasks = []string{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'delete' or 'list' subcommands")
		return
	}

	switch os.Args[1] {
	case "add":
		addTask(os.Args[2:])
	case "delete":
		deleteTask(os.Args[2:])
	case "list":
		listTasks()
	default:
		fmt.Println("Expected 'add', 'delete' or 'list' subcommands")
	}
}

func addTask(arguments []string) {
	if len(arguments) < 1 {
		fmt.Println("Please specify a task to add.")
		return
	}
	task := strings.Join(arguments, " ")
	tasks = append(tasks, task)
	fmt.Printf("Added task: %s\n", task)
}

func deleteTask(arguments []string) {
	if len(arguments) < 1 {
		fmt.Println("Please specify a task number to delete.")
		return
	}
	index := -1
	if len(tasks) == 0 {
		fmt.Println("No tasks to delete.")
		return
	}

	for i, _ := range tasks {
		if fmt.Sprintf("%d", i+1) == arguments[0] {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found.")
		return
	}

	deletedTask := tasks[index]
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Printf("Deleted task: %s\n", deletedTask)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}