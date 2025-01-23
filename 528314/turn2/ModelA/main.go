package main

import (
	"fmt"
	"time"
)

type Task struct {
	TaskName string
	Priority int
	Deadline time.Time
}

// updatePriority adjusts a task's priority based on how close the deadline is.
// For simplicity, we'll increase the priority significantly if the deadline is within 2 days.
func updatePriority(task Task) int {
	now := time.Now()
	daysUntilDeadline := task.Deadline.Sub(now).Hours() / 24

	if daysUntilDeadline <= 2 {
		// Increase priority if the task is due within the next 2 days
		return 5
	}
	return task.Priority
}

// InsertTask inserts a new task into the slice while maintaining the order by descending priority.
func InsertTask(tasks []Task, newTask Task) []Task {
	newTask.Priority = updatePriority(newTask)

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
	// Define tasks with deadlines
	tasks := []Task{
		{"Task 1", 5, time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC)},
		{"Task 2", 3, time.Date(2025, time.February, 10, 0, 0, 0, 0, time.UTC)},
		{"Task 3", 2, time.Date(2025, time.January, 25, 0, 0, 0, 0, time.UTC)},
	}

	fmt.Println("Initial Tasks:", tasks)

	// Insert new tasks
	newTask := Task{"Task 4", 4, time.Date(2025, time.January, 19, 0, 0, 0, 0, time.UTC)}
	tasks = InsertTask(tasks, newTask)
	fmt.Println("After inserting Task 4:", tasks)

	newTask = Task{"Task 5", 1, time.Date(2025, time.January, 16, 0, 0, 0, 0, time.UTC)}
	tasks = InsertTask(tasks, newTask)
	fmt.Println("After inserting Task 5:", tasks)

	newTask = Task{"Task 6", 2, time.Date(2025, time.January, 13, 0, 0, 0, 0, time.UTC)}
	tasks = InsertTask(tasks, newTask)
	fmt.Println("After inserting Task 6:", tasks)
}