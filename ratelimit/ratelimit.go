package ratelimit

import (
	"sync"
	"time"
)

type RateLimiter struct {
	once   sync.Once
	done   chan struct{}
	bucket chan struct{}
}

func NewRateLimiter(limit int, refillRate time.Duration) *RateLimiter {
	// 1. Create buffered channel of size 'limit'
	// 2. Pre-fill the bucket (optional, or let the ticker do it)
	// 3. Start the background ticker
	ticker := time.NewTicker(refillRate)
	rl := RateLimiter{
		bucket: make(chan struct{}, limit),
		done:   make(chan struct{}),
	}

	go func(rateLimiter *RateLimiter, limit int) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
                case rl.bucket <- struct{}{}:
                default:
                }
			case <-rateLimiter.done:
				return
			}
		}
	}(&rl, limit)

	return &rl
}

func (r *RateLimiter) Shutdown() {
	r.once.Do(func() {
		close(r.done)
	})
}

func (r *RateLimiter) Allow() bool {
	select {
    case <-r.bucket:
        return true // Token acquired
    default:
        return false // Bucket empty, return immediately
    }
}
