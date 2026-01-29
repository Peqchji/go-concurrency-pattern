package semaphore

import (
	"fmt"
	"sync"
	"time"
)

// Crawler manages the fetching logic
type Crawler struct {
    semaphore chan struct{}
}

func NewCrawler(limit int) *Crawler {
    return &Crawler{
		semaphore: make(chan struct{}, limit),
	}
}

func (c *Crawler) Crawl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() { <-c.semaphore }()

	// 1. Acquire Token (Block if full)
    // TODO: Send to channel
	c.semaphore <- struct{}{}
	fmt.Printf("🚀 Fetching %s...\n", url)
	time.Sleep(500 * time.Millisecond) // Simulate HTTP Latency

	// 2. Release Token
    // TODO: Read from channel
    
    fmt.Printf("✅ Done %s\n", url)
}

func (c *Crawler) Run() {
	// We want to process 10 URLs, but only allow 3 concurrent fetches
	crawler := NewCrawler(3)
	var wg sync.WaitGroup

	urls := []string{
		"agoda.com", 
		"booking.com", 
		"google.com", 
		"lmwn.com", 
		"expedia.com", 
		"airbnb.com", 
		"trip.com", 
		"kayak.com", 
		"hotels.com", 
		"skyscanner.com",
	}

	for _, u := range urls {
		wg.Add(1)
		// Launch ALL 10 goroutines immediately!
		// The semaphore inside 'Crawl' will stop them from executing all at once.
		go crawler.Crawl(u, &wg)
	}

	wg.Wait()
    fmt.Println("All crawling finished.")
}