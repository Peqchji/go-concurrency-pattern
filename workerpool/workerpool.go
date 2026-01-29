package workerpool

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	wg     sync.WaitGroup
	once   sync.Once

	maxConcurrent int
	jobChan       chan func()
}

func NewWorkerPool(workerNumber int) *WorkerPool {
	return &WorkerPool{
		maxConcurrent: workerNumber,
		jobChan:       make(chan func(), workerNumber),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < int(wp.maxConcurrent); i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for task := range wp.jobChan {
				task()
			}
		}()
	}
}

func (wp *WorkerPool) Submit(task func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("WorkerPool: Submit: panic: %v", r)
		} else {
			err = nil
		}
	}()

	wp.jobChan <- task

	return nil
}

func (wp *WorkerPool) Shutdown() {
	wp.once.Do(func() {
		close(wp.jobChan)
	})

	wp.wg.Wait()
}
