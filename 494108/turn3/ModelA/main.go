package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxWorkers = 5
	rateLimit  = 2 // tasks per second
	taskTimeout = 5 * time.Second
)

type workerPool struct {
	workers        int
	wg             sync.WaitGroup
	ctx            context.Context
	cancelFunc     context.CancelFunc
	workChan       chan<- int
	resultsChan   <-chan int
	rateLimiter    chan struct{}
	completedTasks atomic.Uint64
}

func newWorkerPool(ctx context.Context) *workerPool {
	ctx, cancel := context.WithCancel(ctx)
	rateLimiter := make(chan struct{}, rateLimit)

	return &workerPool{
		ctx:            ctx,
		cancelFunc:     cancel,
		workChan:       make(chan int, 100),
		resultsChan:   make(chan int, 100),
		rateLimiter:    rateLimiter,
		completedTasks: atomic.Uint64(0),
	}
}

func (wp *workerPool) start() {
	for i := 0; i < maxWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *workerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			fmt.Printf("Worker %d received cancellation signal, exiting...\n", id)
			return
		case <-wp.rateLimiter:
			defer func() { wp.rateLimiter <- struct{}{} }()
			select {
			case <-wp.ctx.Done():
				fmt.Printf("Worker %d received cancellation signal during task, exiting...\n", id)
				return
			case work, ok := <-wp.workChan:
				if !ok {
					// workChan closed, no more work to do
					return
				}
				if err := wp.processTask(ctx.WithTimeout(wp.ctx, taskTimeout), work); err != nil {
					fmt.Printf("Worker %d failed to process task %d: %v\n", id, work, err)
				}
			}
		}
	}
}

func (wp *workerPool) processTask(ctx context.Context, work int) error {
	fmt.Printf("Starting task %d...\n", work)
	time.Sleep(time.Duration(work) * time.Second)
	wp.resultsChan <- work
	atomic.AddUint64(&wp.completedTasks, 1)
	fmt.Printf("Task %d completed\n", work)
	return nil
}

func (wp *workerPool) submitWork(work int) {
	select {
	case <-wp.ctx.Done():
		fmt.Printf("Cannot submit work %d: pool is canceled\n", work)
	case wp.workChan <- work:
		fmt.Printf("Work %d submitted\n", work)
	}
}

func (wp *workerPool) trackProgress() {
	for {
		select {
		case result := <-wp.resultsChan:
			fmt.Printf("Received result: %d\n", result)
		case <-wp.ctx.Done():
			fmt.Printf("Progress tracking stopped: pool is canceled\n")
			return
		}
	}
}

func (wp *workerPool) waitForCompletion() {
	fmt.Println("Main thread waiting for workers to finish...")
	wp.wg.Wait()
	fmt.Printf("All workers have finished. Total tasks completed: %d\n", wp.completedTasks.Load())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wp := newWorkerPool(ctx)
	wp.start()
	wp.trackProgress()

	go func() {
		for i := 1; i <= 100; i++ {
			time.Sleep(100 * time.Millisecond)
			wp.submitWork(i)
		}
		close(wp.workChan) // Signal that no more work will be added
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Canceling worker pool...")
	wp.cancelFunc()

	wp.waitForCompletion()
}
