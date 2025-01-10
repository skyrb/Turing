package main  
import (  
    "fmt"
    "sort"
    "time"
)
// Task represents a single task in the to-do list
type Task struct {
    Name    string
    DueDate time.Time
}
// ToDoList is a collection of tasks
type ToDoList struct {
    tasks []Task
}
// NewToDoList creates a new to-do list
func NewToDoList() *ToDoList {  
    return &ToDoList{tasks: make([]Task, 0)}
}
// AddTask adds a new task to the to-do list with a given due date
func (t *ToDoList) AddTask(name string, dueDate time.Time) {
    task := Task{Name: name, DueDate: dueDate}
    t.tasks = append(t.tasks, task)
    sort.Slice(t.tasks, func(i, j int) bool {
        return t.tasks[i].DueDate.Before(t.tasks[j].DueDate)
    })
    fmt.Println("Task added:", name, "due on", dueDate.Format("2006-01-02"))
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
// UpdateTaskDueDate updates the due date of a task by index
func (t *ToDoList) UpdateTaskDueDate(index int, newDueDate time.Time) {
    if index < 0 || index >= len(t.tasks) {
        fmt.Println("Invalid index")
        return
    }
    t.tasks[index].DueDate = newDueDate
    sort.Slice(t.tasks, func(i, j int) bool {
        return t.tasks[i].DueDate.Before(t.tasks[j].DueDate)
    })
    fmt.Println("Task due date updated at index:", index, "new due date:", newDueDate.Format("2006-01-02"))
}
// DisplayList displays all tasks in the to-do list
func (t *ToDoList) DisplayList() {
    if len(t.tasks) == 0 {
        fmt.Println("No tasks in the list")
        return
    }
    for index, task := range t.tasks {
        fmt.Printf("%d. %s (due on %s)\n", index+1, task.Name, task.DueDate.Format("2006-01-02"))
    }
}

func main() {
    todoList := NewToDoList()

    // Adding tasks with due dates
    todoList.AddTask("Buy groceries", time.Date(2023, 7, 15, 0, 0, 0, 0, time.UTC))
    todoList.AddTask("Finish development", time.Date(2023, 6, 30, 0, 0, 0, 0, time.UTC))
    todoList.AddTask("Exercise", time.Date(2023, 7, 20, 0, 0, 0, 0, time.UTC))
    
    todoList.DisplayList()

    // Updating due date
    todoList.UpdateTaskDueDate(1, time.Date(2023, 7, 25, 0, 0, 0, 0, time.UTC))

    todoList.DisplayList()

    // Deleting a task
    todoList.DeleteTask(0)

    todoList.DisplayList()
}