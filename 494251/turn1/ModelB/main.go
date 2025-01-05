package main  
import (  
    "fmt"
    "sync"
)

type ReadWriteQueue interface {  
    Read() string
    Write(string)
    Size() int
}

type ConcurrentQueue struct {  
    mutex sync.RWMutex
    queue []string
}  

func NewConcurrentQueue() ReadWriteQueue {  
    return &ConcurrentQueue{}
}

func (q *ConcurrentQueue) Read() string {  
    q.mutex.RLock()
    defer q.mutex.RUnlock()

    if len(q.queue) == 0 {  
        return ""
    }
    return q.queue[0]
}

func (q *ConcurrentQueue) Write(item string) {  
    q.mutex.Lock()
    defer q.mutex.Unlock()
    q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Size() int {  
    q.mutex.RLock()
    defer q.mutex.RUnlock()
    return len(q.queue)
}

func main() {  
    q := NewConcurrentQueue()

    //Producers
    go func() {
        for i := 0; i < 10; i++ {
            q.Write(fmt.Sprintf("File-%d", i))
        }
    }()

    //Consumers
    go func() {
        for q.Size() > 0 {
            fmt.Println("Reading:", q.Read())
        }
    }()

    // main thread will keep running until both producer and consumer are done
}  
