package atomic
// The Concept: For simple counters (metrics, request IDs), sync.Mutex is heavy (locks the OS thread/goroutine). sync/atomic uses 
// CPU-level instructions (CAS - Compare And Swap) which are much faster for simple arithmetic.

// The Task:
// - Create a struct ServerMetrics with a field requestCount int64.
// - Spawn 100 goroutines that increment this counter 1000 times each.
// - Implement version A using sync.Mutex.
// - Implement version B using atomic.AddInt64.
// - (Optional) Benchmark them using Go's testing tool, or just observe the code complexity difference.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ServerMetrics struct {
	// mu           sync.Mutex
	requestCount int64
}

func NewServerMetrics() *ServerMetrics {
	return &ServerMetrics{
		// mu: sync.Mutex{},
		requestCount: 0,
	}
}

func (sm *ServerMetrics) IncreseRequestCount(value int64) {
	atomic.AddInt64(&sm.requestCount, value)
}

func task(metrics *ServerMetrics) {
	for i := 0; i < 1000; i += 1 {
		metrics.IncreseRequestCount(1)
	}
}

func Atomic() {
	var wg sync.WaitGroup
	metric := NewServerMetrics()

	for i := 0; i < 100; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			task(metric)
		}()
	}

	wg.Wait()
	fmt.Println(metric.requestCount)
}