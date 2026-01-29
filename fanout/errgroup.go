package fanout

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetchData(ctx context.Context, id int) error {
	select {
	case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
		if id == 3 { // Simulate failure on ID 3
			return fmt.Errorf("fetcher %d failed", id)
		}
		fmt.Printf("Fetcher %d success\n", id)
		return nil
	case <-ctx.Done():
		fmt.Printf("Fetcher %d canceled\n", id)
		return ctx.Err()
	}
}

func FanOutWithErrGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 5; i++ {
		id := i
		g.Go(func() error {
			return fetchData(ctx, id)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Batch failed: %v\n", err)
	} else {
		fmt.Println("Batch success!")
	}
}