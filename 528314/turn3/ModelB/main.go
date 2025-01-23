package main  
import (  
    "container/heap"
    "fmt"
    "sync"
    "time"
)

type Task struct {  
    TaskName string
    Priority int
    Deadline time.Time
}  

type taskHeap []Task  

func (h taskHeap) Len() int           { return len(h) }
func (h taskHeap) Less(i, j int) bool { return h[i].Priority > h[j].Priority }
func (h taskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *taskHeap) Push(x interface{}) {  
    *h = append(*h, x.(Task))
}

func (h *taskHeap) Pop() interface{} {  
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}  

type taskManager struct {  
    tasks taskHeap  
    mu    sync.Mutex
}  

func (tm *taskManager) InsertTask(task Task) {  
    tm.mu.Lock()
    defer tm.mu.Unlock()

    heap.Push(&tm.tasks, task)
}  

func (tm *taskManager) UpdateTaskPriority(taskName string, newPriority int) {  
    tm.mu.Lock()
    defer tm.mu.Unlock()

    for i, task := range tm.tasks {
        if task.TaskName == taskName {
            tm.tasks[i].Priority = newPriority
            heap.Fix(&tm.tasks, i)
            return
        }
    }
    fmt.Println("Task not found:", taskName)
}  

func (tm *taskManager) CompleteTask(taskName string) {  
    tm.mu.Lock()
    defer tm.mu.Unlock()

    for i, task := range tm.tasks {
        if task.TaskName == taskName {
            tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
            heap.Fix(&tm.tasks, i) // Adjust heap after removal
            return
        }
    }
    fmt.Println("Task not found:", taskName)
}  

func (tm *taskManager) TopTask() (Task, bool) {  
    tm.mu.Lock()
    defer tm.mu.Unlock()

    if len(tm.tasks) == 0 {
        return Task{}, false
    }
    return tm.tasks[0], true
}  

func main() {  
    tm := &taskManager{}

    // Insert tasks
    tm.InsertTask(Task{"Task 3", 2, time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC)})  
    tm.InsertTask(Task{"Task 1", 5, time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)})
    tm.InsertTask(Task{"Task 2", 3, time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)})
  
    topTask, ok := tm.TopTask()
    if ok {
        fmt.Println("Top Task:", topTask.TaskName, "Priority:", topTask.Priority)