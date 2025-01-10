package main  
import (  
    "fmt"
    "sort"
    "time"
)

// TaskRecurrence represents the recurrence type of a task
type TaskRecurrence int

const (
    NoRecurrence TaskRecurrence = iota
    Daily
    Weekly
    Monthly
)

func (r TaskRecurrence) String() string {
    switch r {
    case Daily:
        return "Daily"
    case Weekly:
        return "Weekly"
    case Monthly:
        return "Monthly"
    default:
        return "No Recurrence"
    }
}

// Task represents a task with a description, deadline, recurrence, and completion status
type Task struct {
    Description string
    Deadline    time.Time
    Recurrence  TaskRecurrence
    Completed   bool
}

// ToDoList is a collection of tasks
type ToDoList struct {
    tasks []Task
}

// NewToDoList creates a new to-do list
func NewToDoList() *ToDoList {
    return &ToDoList{tasks: make([]Task, 0)}
}

// AddTask adds a new task with a deadline and recurrence to the to-do list
func (t *ToDoList) AddTask(description string, deadline time.Time, recurrence TaskRecurrence) {
    newTask := Task{Description: description, Deadline: deadline, Recurrence: recurrence}
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

// MarkTaskComplete marks a task as complete at the given index
func (t *ToDoList) MarkTaskComplete(index int) {
    if index < 0 || index >= len(t.tasks) {
        fmt.Println("Invalid index")
        return
    }
    t.tasks[index].Completed = true
    fmt.Println("Task marked as complete at index:", index)
}

// DisplayList displays all tasks in the to-do list
func (t *ToDoList) DisplayList() {
    if len(t.tasks) == 0 {
        fmt.Println("No tasks in the list")
        return
    }
    for index, task := range t.tasks {
        completedStatus := "Incomplete"
        if task.Completed {
            completedStatus = "Completed"
        }
        fmt.Printf("%d. %s (Due: %s, Recurrence: %s, Status: %s)\n", index+1, task.Description, task.Deadline.Format("2006-01-02"), task.Recurrence, completedStatus)
    }
}

// SortByDeadline sorts the tasks by their deadlines
func (t *ToDoList) SortByDeadline() {
    sort.Slice(t.tasks, func(i, j int) bool {
        return t.tasks[i].Deadline.Before(t.tasks[j].Deadline)
    })