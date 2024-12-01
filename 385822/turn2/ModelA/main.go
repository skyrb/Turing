package main  
import (
    "context"
    "sync"
    "time"
)

type Task struct {
    Name     string   // A name for the task
    Run      func(context.Context) // The function to run
    Duration time.Duration // Estimated execution duration
    timeout  time.Duration // Task timeout
}

type Scheduler struct {
    tasks chan Task
    wg    sync.WaitGroup
}

func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make(chan Task),
    }
}

func (s *Scheduler) Start() {
    s.wg.Add(1)
    go s.run()
}

func (s *Scheduler) Stop() {
    close(s.tasks)
    s.wg.Wait()
}

// Submit a new task to the scheduler.
func (s *Scheduler) Submit(t Task) {
    s.tasks <- t
}

// The main scheduling loop
func (s *Scheduler) run() {
    for task := range s.tasks {
        ctx, cancel := context.WithTimeout(context.Background(), task.timeout)
        defer cancel()

        s.wg.Add(1)
        go func(t Task, ctx context.Context) {
            defer s.wg.Done()
            select {
            case <-ctx.Done():
                println("Task", t.Name, "cancelled")
                return
            default:
            }
            time.Sleep(t.Duration)
            println("Ran task:", t.Name, "for", t.Duration)
        }(task, ctx)
    }
    s.wg.Done()
}

func longRunningTask(ctx context.Context) {
    time.Sleep(5 * time.Second)
}

func main() {
    scheduler := NewScheduler()
    scheduler.Start()

    // Task with timeout
    scheduler.Submit(Task{
        Name:     "Task C",
        Run:      longRunningTask,
        Duration: 2 * time.Second,
        timeout:  3 * time.Second,
    })

    // Task without timeout
    scheduler.Submit(Task{
        Name:     "Task A",
        Run:      longRunningTask,
        Duration: 1 * time.Second,
    })

    // Task with quick cancellation
    scheduler.Submit(Task{
        Name:     "Task B",
        Run:      longRunningTask,
        Duration: 4 * time.Second,
        timeout:  100 * time.Millisecond,
    })

    time.Sleep(6 * time.Second) // Wait for tasks to complete or timeout
    scheduler.Stop()
    println("Scheduler stopped")
}