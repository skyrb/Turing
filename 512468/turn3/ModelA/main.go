package main

import (
	"fmt"
	"sort"
	"time"
)

// Task represents a non-recurring task with a description and a deadline
type Task struct {
	Description string
	Deadline    time.Time
	Completed   bool
}

// RecurringTask represents a recurring task with a description, frequency, and next deadline
type RecurringTask struct {
	Description string
	Frequency   string // e.g., "Daily", "Weekly", "Monthly"
	NextDeadline time.Time
	Completed   bool
}

// ToDoList is a collection of tasks and recurring tasks
type ToDoList struct {
	tasks       []Task
	recurring   []RecurringTask
}

// NewToDoList creates a new to-do list
func NewToDoList() *ToDoList {
	return &ToDoList{tasks: make([]Task, 0), recurring: make([]RecurringTask, 0)}
}

// AddTask adds a new non-recurring task with a deadline to the to-do list
func (t *ToDoList) AddTask(description string, deadline time.Time) {
	newTask := Task{Description: description, Deadline: deadline, Completed: false}
	t.tasks = append(t.tasks, newTask)
	fmt.Println("Task added:", newTask)
}

// UpdateTaskDeadline updates the deadline for a non-recurring task at the given index
func (t *ToDoList) UpdateTaskDeadline(index int, newDeadline time.Time) {
	if index < 0 || index >= len(t.tasks) {
		fmt.Println("Invalid index")
		return
	}
	t.tasks[index].Deadline = newDeadline
	fmt.Println("Deadline updated for task at index:", index)
}

// DeleteTask deletes a non-recurring task from the to-do list by index
func (t *ToDoList) DeleteTask(index int) {
	if index < 0 || index >= len(t.tasks) {
		fmt.Println("Invalid index")
		return
	}
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
	fmt.Println("Task deleted at index:", index)
}

// DisplayNonRecurringList displays all non-recurring tasks in the to-do list
func (t *ToDoList) DisplayNonRecurringList() {
	if len(t.tasks) == 0 {
		fmt.Println("No non-recurring tasks in the list")
		return
	}
	for index, task := range t.tasks {
		fmt.Printf("%d. %s (Due: %s) - Completed: %v\n", index+1, task.Description, task.Deadline.Format("2006-01-02"), task.Completed)
	}
}

// AddRecurringTask adds a new recurring task with a frequency and next deadline to the to-do list
func (t *ToDoList) AddRecurringTask(description string, frequency string, nextDeadline time.Time) {
	newRecurring := RecurringTask{Description: description, Frequency: frequency, NextDeadline: nextDeadline, Completed: false}
	t.recurring = append(t.recurring, newRecurring)
	fmt.Println("Recurring task added:", newRecurring)
}

// UpdateRecurringTaskDeadline updates the next deadline for a recurring task at the given index
func (t *ToDoList) UpdateRecurringTaskDeadline(index int, newDeadline time.Time) {
	if index < 0 || index >= len(t.recurring) {
		fmt.Println("Invalid index")
		return
	}
	t.recurring[index].NextDeadline = newDeadline
	fmt.Println("Next deadline updated for recurring task at index:", index)
}

// DeleteRecurringTask deletes a recurring task from the to-do list by index
func (t *ToDoList) DeleteRecurringTask(index int) {
	if index < 0 || index >= len(t.recurring) {
		fmt.Println("Invalid index")
		return
	}
	t.recurring = append(t.recurring[:index], t.recurring[index+1:]...)
	fmt.Println("Recurring task deleted at index:", index)
}

// DisplayRecurringList displays all recurring tasks in the to-do list
func (t *ToDoList) DisplayRecurringList() {
	if len(t.recurring) == 0 {
		fmt.Println("No recurring tasks in the list")
		return
	}
	for index, task := range t.recurring {
		fmt.Printf("%d. %s (Frequency: %s, Next Due: %s) - Completed: %v\n", index+1, task.Description, task.Frequency, task.NextDeadline.Format("2006-01-02"), task.Completed)
	}
}

// SortByDeadline sorts the tasks and recurring tasks by their deadlines
func (t *ToDoList) SortByDeadline() {
	sort.Slice(t.tasks, func(i, j int) bool {
		return t.tasks[i].Deadline.Before(t.tasks[j].Deadline)
	})
	sort.Slice(t.recurring, func(i, j int) bool {
		return t.recurring[i].NextDeadline.Before(t.recurring[j].NextDeadline)
	})
}

func main() {
	todoList := NewToDoList()

	// Add non-recurring tasks with deadlines
	todoList.AddTask("Buy groceries", time.Now().AddDate(0, 0, 2)) // Due in 2 days
	todoList.AddTask("Finish development", time.Now().AddDate(0, 0, 5)) // Due in 5 days
	todoList.AddTask("Exercise", time.Now().AddDate(0, 0, 1)) // Due in 1 day

	// Add recurring tasks with frequencies and next deadlines
	todoList.AddRecurringTask("Daily workout", "Daily", time.Now().AddDate(0, 0, 1)) // Next due in 1 day
	todoList.AddRecurringTask("Weekly meeting", "Weekly", time.Now().AddDate(0, 0, 7)) // Next due in 7 days
	todoList.AddRecurringTask("Monthly review", "Monthly", time.Now().AddDate(0, 1, 0)) // Next due in 1 month

	// Display lists
	fmt.Println("\nNon-recurring tasks:")
	todoList.DisplayNonRecurringList()

	fmt.Println("\nRecurring tasks:")
	todoList.DisplayRecurringList()

	// Sort by deadline
	todoList.SortByDeadline()
	fmt.Println("\nSorted by deadline:")

	fmt.Println("\nNon-recurring tasks:")
	todoList.DisplayNonRecurringList()

	fmt.Println("\nRecurring tasks:")
	todoList.DisplayRecurringList()

	// Update deadlines
	todoList.UpdateTaskDeadline(1, time.Now().AddDate(0, 0, 3)) // Update "Finish development" to due in 3 days
	todoList.UpdateRecurringTaskDeadline(2, time.Now().AddDate(0, 0, 14)) // Update "Weekly meeting" to next due in 14 days

	// Display lists after updating
	fmt.Println("\nAfter updating deadlines:")

	fmt.Println("\nNon-recurring tasks:")
	todoList.DisplayNonRecurringList()

	fmt.Println("\nRecurring tasks:")
	todoList.DisplayRecurringList()

	// Delete tasks
	todoList.DeleteTask(0) // Delete "Buy groceries"
	todoList.DeleteRecurringTask(1) // Delete "Daily workout"

	// Display lists after deletion
	fmt.Println("\nAfter deleting tasks:")

	fmt.Println("\nNon-recurring tasks:")
	todoList.DisplayNonRecurringList()

	fmt.Println("\nRecurring tasks:")
	todoList.DisplayRecurringList()
}