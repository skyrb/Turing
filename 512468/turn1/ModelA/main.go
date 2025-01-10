package main

import (
	"fmt"
)

// Task represents a single task in the to-do list
type Task string

// ToDoList is a collection of tasks
type ToDoList struct {
	tasks []Task
}

// NewToDoList creates a new to-do list
func NewToDoList() *ToDoList {
	return &ToDoList{tasks: make([]Task, 0)}
}

// AddTask adds a new task to the to-do list
func (t *ToDoList) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
	fmt.Println("Task added:", task)
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
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func main() {
	todoList := NewToDoList()

	todoList.AddTask("Buy groceries")
	todoList.AddTask("Finish development")
	todoList.AddTask("Exercise")

	todoList.DisplayList()

	todoList.DeleteTask(1)
	todoList.DisplayList()

	todoList.DeleteTask(10) // Invalid index
}