package main
import (
    "sync"
    "time"
)

type Task struct {
    Name     string   // A name for the task
    Run      func()   // The function to run
    Duration time.Duration // Estimated execution duration
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
        // In a real implementation, you'd use a more sophisticated priority queue.
        // For simplicity, here we just select the task with the shortest duration.

        //var currentTask Task
        //for currentTask == (Task{}) {
        //    select {
        //    case currentTask = <-s.tasks:
        //    default:
        //    }
        //}


        s.wg.Add(1)
        go func(t Task) {
            defer s.wg.Done()
            time.Sleep(t.Duration)
            println("Ran task:", t.Name, "for", t.Duration)
        }(task)
    }
    s.wg.Done()
}

func main() {
    scheduler := NewScheduler()
    scheduler.Start()

    // Submit tasks to the scheduler
    scheduler.Submit(Task{
        Name:     "Task C",
        Run:      func() {},
        Duration: 2 * time.Second,
    })

    scheduler.Submit(Task{
        Name:     "Task A",
        Run:      func() {},
        Duration: 1 * time.Second,
    })

    scheduler.Submit(Task{
        Name:     "Task B",
        Run:      func() {},
        Duration: 3 * time.Second,
    })

    time.Sleep(4 * time.Second) // Wait for tasks to complete
    scheduler.Stop()
    println("Scheduler stopped")
}