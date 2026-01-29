package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var replicaA = map[string]string{ "id_1": "Alice", "id_3": "Charlie", "id_5": "Eve (A)" }
var replicaB = map[string]string{ "id_2": "Bob", "id_4": "Dave", "id_5": "Eve (B)" }

// Search now respects Context and Latency
func Search(ctx context.Context, replicaID int, target string) (string, error) {
	// 1. Simulate Network Latency (Critical to see concurrency working)
	select {
	case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
		// Continue search
	case <-ctx.Done():
		// Stop immediately if another worker won
		return "", ctx.Err()
	}

	// 2. Perform Search
	var db map[string]string
	if replicaID == 0 { db = replicaA } else { db = replicaB }

	if val, exists := db[target]; exists {
		return val, nil
	}
	return "", fmt.Errorf("not found")
}

func ParallelSearch() {
	// We use WithCancel so we can kill the loser
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Safety: Ensure context is cleaned up

	// Buffered channel to prevent blocking if we exit early
	resultChan := make(chan string, 2)

	// Launch Workers
	for id := 0; id < 2; id++ {
		// FIX: Pass 'id' as argument to prevent loop variable capture
		go func(workerID int) {
			foundName, err := Search(ctx, workerID, "id_2")
			
			if err == nil {
				// Found it! Send success
				resultChan <- foundName
			} else {
				// Failed. Send empty string to signal "I'm done but found nothing"
				resultChan <- ""
			}
		}(id)
	}

	// Wait for First Success OR All Failures
	// We need to count failures to know when to give up
	failureCount := 0
	for i := 0; i < 2; i++ {
		res := <-resultChan
		if res != "" {
			fmt.Println("🎉 Found:", res)
			cancel() // STOP the other worker immediately!
			return
		}
		failureCount++
	}

	fmt.Println("❌ Not found in any replica.")
}
