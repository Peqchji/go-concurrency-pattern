package singleflight

import (
	"sync"
	"time"
)

// call is an in-flight or completed singleflight.do call
type Call[T any] struct {
	once   sync.Once
	result T
	err    error
}

func (c *Call[T]) Do(fn func() (T, error)) (T, error) {
	c.once.Do(func() {
		time.Sleep(3 * time.Second)
		c.result, c.err = fn() 
	})

	return c.result, c.err
}

// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group struct {
	// TODO: You need a lock to protect the map of active calls
	// TODO: You need a map to track keys currently "in-flight"
	mu     sync.RWMutex
	memory map[string]*Call[any]
}

// NewGroup creates a new Singleflight Group
func NewGroup() *Group {
	return &Group{
		memory: make(map[string]*Call[any]),
	}
}

// Do executes and returns the results of the given function, making sure that
// only one execution is in-flight for a given key at a time.
//
//  1. If a duplicate comes in, the duplicate caller waits for the original to complete
//     and receives the same results.
//  2. The return value shared indicates whether v was given to multiple callers.
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
    // 1. Read Path (Fast)
    g.mu.RLock()
    if caller, exists := g.memory[key]; exists {
        g.mu.RUnlock()
        res, err := caller.Do(fn) // Wait for the existing job
        return res, err, true
    }
    g.mu.RUnlock()

    // 2. Write Path (Slow)
    g.mu.Lock()
    
    // 3. Double Check
    if caller, exists := g.memory[key]; exists {
        g.mu.Unlock()             // Release lock immediately!
        res, err := caller.Do(fn) // Wait
        return res, err, true
    }

    // 4. Register new job
    caller := &Call[any]{}
    g.memory[key] = caller
    g.mu.Unlock() // <--- CRITICAL: Release lock BEFORE running the slow job!

    // 5. Execute Job (No Map Lock held here!)
    res, err := caller.Do(fn)

    // 6. Cleanup (Re-acquire Lock)
    g.mu.Lock()
    delete(g.memory, key)
    g.mu.Unlock()

    return res, err, false
}