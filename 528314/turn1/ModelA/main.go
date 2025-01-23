package main

import (
	"fmt"
)

type Task struct {
	TaskName string
	Priority int
}

// InsertTask inserts a new task into the slice while maintaining the order by descending priority.
func InsertTask(tasks []Task, newTask Task) []Task {
	// Find the correct position to insert the new task
	position := 0
	for i, task := range tasks {
		if newTask.Priority > task.Priority {
			position = i
			break
		}
		position = i + 1
	}

	// Insert the new task at the found position
	tasks = append(tasks[:position], append([]Task{newTask}, tasks[position:]...)...)

	return tasks
}

func main() {
	// Initial tasks
	tasks := []Task{
		{"Task 1", 5},
		{"Task 2", 3},
		{"Task 3", 1},
	}

	fmt.Println("Initial Tasks:", tasks)

	// Insert new tasks
	newTask := Task{"Task 4", 4}
	tasks = InsertTask(tasks, newTask)
	fmt.Println("After inserting Task 4:", tasks)

	newTask = Task{"Task 5", 6}
	tasks = InsertTask(tasks, newTask)
	fmt.Println("After inserting Task 5:", tasks)

	newTask = Task{"Task 6", 2}
	tasks = InsertTask(tasks, newTask)
	fmt.Println("After inserting Task 6:", tasks)
}