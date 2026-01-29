package fan_in

import (
	"fmt"
	"sync"
	"time"
)

// The Concept: "Fan-In"
// You have a pile of Jobs (e.g., 100 images to resize). You have a team of Workers (e.g., 3 goroutines).
// You want One Result Pipe where all the finished work ends up.

// The Trap: If you have 3 workers writing to results channel, who closes the channel?

// If Worker 1 closes it, Worker 2 will panic when it tries to write.

// If Main closes it, it might close it while workers are still working (Panic).

// Solution: You need a "Foreman" (a separate goroutine) who waits for everyone to punch out, then closes the door.

// 1. The Worker
// Reads numbers from 'jobs', multiplies by 2, sleeps 10ms, sends to 'results'.
// Uses WaitGroup to signal "I am done".
type WorkerPool[T any] struct {
	wg       	sync.WaitGroup
	once     	sync.Once
	workerCount int

	resultsChan chan T
	jobsChan 	chan func() T
}

func NewWorkerPool[T any](workerCount, queueSize int) *WorkerPool[T] {
	pool := &WorkerPool[T]{
		workerCount: workerCount,
		jobsChan: make(chan func() T, queueSize),
		resultsChan: make(chan T, queueSize),
	}

	return pool
}

func (wp *WorkerPool[T]) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker executes tasks from the queue
func (wp *WorkerPool[T]) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.jobsChan {
		fmt.Printf("Worker %d executing task...\n", id)
		time.Sleep(time.Second)
		wp.resultsChan <- task()
	}
}

func (wp *WorkerPool[T]) Consume() <-chan T {
	return wp.resultsChan
}

func (wp *WorkerPool[T]) Submit(job func() T) {
	wp.jobsChan <- job
}

func (wp *WorkerPool[T]) Shutdown() {
	wp.once.Do(func(){
		close(wp.jobsChan)

		go func() {
            wp.wg.Wait()          // Wait for all active tasks to finish
            close(wp.resultsChan) // Safe to close now
        }()
	})
}

func FanIn() {
    pool := NewWorkerPool[string](10, 10)
    pool.Start()

    go func() {
        for i := 1; i <= 1_000; i++ {
            id := i
            pool.Submit(func() string {
                return fmt.Sprintf("Processed User ID: %d", id)
            })
        }

        pool.Shutdown()
    }()

    for result := range pool.Consume() {
        fmt.Println("Received:", result)
    }

    fmt.Println("All work completed.")
}
