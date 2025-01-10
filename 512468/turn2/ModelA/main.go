package main

import (
	"fmt"
	"sort"
	"time"
)

// Task represents a task with a description and a deadline
type Task struct {
	Description string
	Deadline    time.Time
}

// ToDoList is a collection of tasks
type ToDoList struct {
	tasks []Task
}

// NewToDoList creates a new to-do list
func NewToDoList() *ToDoList {
	return &ToDoList{tasks: make([]Task, 0)}
}

// AddTask adds a new task with a deadline to the to-do list
func (t *ToDoList) AddTask(description string, deadline time.Time) {
	newTask := Task{Description: description, Deadline: deadline}
	t.tasks = append(t.tasks, newTask)
	fmt.Println("Task added:", newTask)
}

// UpdateTaskDeadline updates the deadline for a task at the given index
func (t *ToDoList) UpdateTaskDeadline(index int, newDeadline time.Time) {
	if index < 0 || index >= len(t.tasks) {
		fmt.Println("Invalid index")
		return
	}
	t.tasks[index].Deadline = newDeadline
	fmt.Println("Deadline updated for task at index:", index)
}

// DeleteTask deletes a task from the to-do list by index
func (t *ToDoList) DeleteTask(index int) {
	if index < 0 || index >= len(t.tasks) {
		fmt.Println("Invalid index")
		return
	}
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
	fmt.Println("Task deleted at index:", index)
}

// DisplayList displays all tasks in the to-do list
func (t *ToDoList) DisplayList() {
	if len(t.tasks) == 0 {
		fmt.Println("No tasks in the list")
		return
	}
	for index, task := range t.tasks {
		fmt.Printf("%d. %s (Due: %s)\n", index+1, task.Description, task.Deadline.Format("2006-01-02"))
	}
}

// SortByDeadline sorts the tasks by their deadlines
func (t *ToDoList) SortByDeadline() {
	sort.Slice(t.tasks, func(i, j int) bool {
		return t.tasks[i].Deadline.Before(t.tasks[j].Deadline)
	})
}

func main() {
	todoList := NewToDoList()

	// Add tasks with deadlines
	todoList.AddTask("Buy groceries", time.Now().AddDate(0, 0, 2)) // Due in 2 days
	todoList.AddTask("Finish development", time.Now().AddDate(0, 0, 5)) // Due in 5 days
	todoList.AddTask("Exercise", time.Now().AddDate(0, 0, 1)) // Due in 1 day

	// Display the list
	todoList.DisplayList()

	// Sort the list by deadline
	todoList.SortByDeadline()
	fmt.Println("\nSorted by deadline:")
	todoList.DisplayList()

	// Update the deadline for a task
	todoList.UpdateTaskDeadline(1, time.Now().AddDate(0, 0, 3)) // Update "Finish development" to due in 3 days

	// Display the list after updating
	todoList.DisplayList()

	// Sort again by deadline
	todoList.SortByDeadline()
	fmt.Println("\nSorted by deadline after update:")
	todoList.DisplayList()

	// Delete a task
	todoList.DeleteTask(0) // Delete "Buy groceries"

	// Display the list after deletion
	todoList.DisplayList()
}