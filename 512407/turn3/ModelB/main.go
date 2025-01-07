package main  
import (  
    "fmt"
    "sync"
    "time"
)

type Worker struct {
    id       int
    wg       *sync.WaitGroup
    jobs     <-chan int
    stopChan chan struct{}
}

func newWorker(id int, wg *sync.WaitGroup, jobs <-chan int, stopChan chan struct{}) *Worker {
    return &Worker{
        id:       id,
        wg:       wg,
        jobs:     jobs,
        stopChan: stopChan,
    }
}

func (w *Worker) run() {
    defer w.wg.Done()
    for {
        select {
        case job := <-w.jobs:
            fmt.Printf("Worker %d started job %d\n", w.id, job)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d finished job %d\n", w.id, job)
        case <-w.stopChan:
            fmt.Printf("Worker %d received stop signal, exiting.\n", w.id)
            return // <= *** This is where the goroutine should end.
        }
    }
}

func main() {
    numWorkers := 3  
    numJobs := 10  
    jobs := make(chan int, numJobs)
    var wg sync.WaitGroup
    
    stopChan := make(chan struct{})

    // Start workers  
    for i := 1; i <= numWorkers; i++ {
        worker := newWorker(i, &wg, jobs, stopChan)
        wg.Add(1)
        go worker.run()
    }  
    
    // Send jobs  
    for i := 1; i <= numJobs; i++ {
        jobs <- i
    }

    // wait for a bit
    time.Sleep(2 * time.Second)

    //** closing stopChannel should trigger workers to stop, but  it doesn't**
    close(stopChan)
    
    wg.Wait()
    fmt.Println("All workers stopped. But where are the busy workers?")

    //  Goroutine count will be higher than expected after 2 seconds
    time.Sleep(5 * time.Second)
    fmt.Println("Program ending...")
}  