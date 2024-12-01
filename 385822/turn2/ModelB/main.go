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

func main() {
    scheduler := NewScheduler()
    scheduler.Start()

    // Task with timeout
    scheduler.Submit(Task{
        Name:     "Task A",
        Run:      func(ctx context.Context) {},
        Duration: 5 * time.Second,
        timeout:  2 * time.Second,
    })

    // Task without timeout
    scheduler.Submit(Task{
        Name:     "Task B",
        Run:      func(ctx context.Context) {},
        Duration: 3 * time.Second,
    })

    // Cancel a task after 1 second
    time.AfterFunc(1*time.Second, func() {
        scheduler.Submit(Task{
            Name:     "Task C",
            Run:      func(ctx context.Context) {},
            Duration: 4 * time.Second,
            timeout:  0, // No timeout, but will be cancelled after 1 sec
        })
    })

    time.Sleep(6 * time.Second) // Wait for tasks to complete or cancel
    scheduler.Stop()
    println("Scheduler stopped")
}