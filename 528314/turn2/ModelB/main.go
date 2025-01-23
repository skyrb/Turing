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

func InsertTask(tasks []Task, newTask Task) []Task {  
    position := 0
    for i, task := range tasks {
        if newTask.Priority > task.Priority {  
            position = i
            break
        }
        position = i + 1
    }
    tasks = append(tasks[:position], append([]Task{newTask}, tasks[position:]...)...)
    return tasks
}

func updatePriorityByDeadline(tasks []Task) []Task {
    today := time.Now().Date()
    for i, task := range tasks {
        daysToDeadline := task.Deadline.Sub(today).Hours() / 24
        if daysToDeadline < 0 {
            tasks[i].Priority = 1 // Set priority to 1 for overdue tasks
        } else if daysToDeadline < 2 {
            tasks[i].Priority = 5 // Set priority to 5 for tasks due in 2 days or less
        } else if daysToDeadline < 5 {
            tasks[i].Priority = 3 // Set priority to 3 for tasks due in 5 days or less
        }
    }
    return tasks
}

func main() {  
    tasks := []Task{  
        {"Task 1", 5, time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)},
        {"Task 2", 3, time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
        {"Task 3", 2, time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC)},  
    }  

    fmt.Println("Initial Tasks:")
    for _, task := range tasks {
        fmt.Printf("%s, Priority: %d, Deadline: %s\n", task.TaskName, task.Priority, task.Deadline.Format("2006-01-02"))
    }

    tasks = updatePriorityByDeadline(tasks)

    fmt.Println("\nTasks after updating priority by deadline:")
    for _, task := range tasks {
        fmt.Printf("%s, Priority: %d, Deadline: %s\n", task.TaskName, task.Priority, task.Deadline.Format("2006-01-02"))
    }

    newTask := Task{"Task 4", 4, time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC)}  
    tasks = InsertTask(tasks, newTask)  
    tasks = updatePriorityByDeadline(tasks)

    fmt.Println("\nTasks after inserting Task 4 and updating priority:")
    for _, task := range tasks {
        fmt.Printf("%s, Priority: %d, Deadline: %s\n", task.TaskName, task.Priority, task.Deadline.Format("2006-01-02"))
    }
}