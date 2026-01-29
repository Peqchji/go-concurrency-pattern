package pipeline

import (
	"fmt"
)

// Stage 1: Generator
// Converts a variadic list of integers into a read-only channel.
// Must close the channel when done.
func gen(nums ...int) <-chan int {
    // TODO: Implement
	out := make(chan int, 100)

	go func() {
		defer close(out)

		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

// Stage 2: Square
// Reads from 'in', squares the number, and writes to a NEW output channel.
// Must close the output channel when 'in' is closed.
func sq(in <-chan int) <-chan int {
    // TODO: Implement
	out := make(chan int, 100)

	go func() {
		defer close(out)
		
		for num := range in {
			out <- num*num
		}
	}()

	return out
}

// Stage 3: Main
// 1. Set up the pipeline: gen(2, 3) -> sq -> print result
// 2. Consume the final output using 'range'.
func Pipeline() {
    // TODO: Implement
	pipeline := gen(2, 3, 100, 10000, 10101010, 1)
	resultChan := sq(pipeline)

	for result := range resultChan {
		fmt.Println("result: ", result)
	}
}