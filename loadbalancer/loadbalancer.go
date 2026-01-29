package loadbalancer

import (
	"sync"
)

// SafeSet wraps a map with a RWMutex
type LoadBalancer struct {
	mu   sync.RWMutex
	currentService int
	serviceRegistry map[int]string
}

// NewSafeSet initializes the map
func NewLoadBalancer(serviceRegistry map[int]string) *LoadBalancer {
	return &LoadBalancer{
		serviceRegistry: serviceRegistry,
	}
}

// Add safely adds a key (Write Lock)
func (lb *LoadBalancer) Next() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	cur := lb.currentService
	lb.currentService = (lb.currentService + 1)%len(lb.serviceRegistry)

	return lb.serviceRegistry[cur]
}
