package threadsafeset

import (
	"sync"
)

// SafeSet wraps a map with a RWMutex
type SafeSet struct {
	mu   sync.RWMutex
	data map[string]struct{}
}

// NewSafeSet initializes the map
func NewSafeSet() *SafeSet {
	return &SafeSet{
		data: make(map[string]struct{}),
	}
}

// Add safely adds a key (Write Lock)
func (s *SafeSet) Add(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = struct{}{}
}

// Contains safely checks a key (Read Lock)
// Optimization: Use RLock() allows multiple readers at once!
func (s *SafeSet) Contains(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	_, isContain := s.data[key]

	return isContain
}
