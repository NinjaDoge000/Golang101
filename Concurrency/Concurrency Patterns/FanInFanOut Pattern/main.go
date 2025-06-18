package main

import (
	"fmt"
	"sync"
	"time"
)

/*

	It's a concurrency pattern used to:
	Fan-Out: Distribute work across multiple goroutines.
	Fan-In: Collect results from those multiple goroutines into a single channel.

	Insight:
	1. Channel safety: All workers share the same input and output channel — safe in Go due to internal locking in channels.
	2. WaitGroup is used to detect when all workers are done, and then close the output channel.
	3. Closing channels: Only close from the sending side, once all sends are done.


*/

func generator(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, val := range nums {
			out <- val
		}
		close(out)
	}()

	return out
}

func worker(in <-chan int, result chan<- int, wg *sync.WaitGroup) {

	go func() {

		defer wg.Done()
		for n := range in {
			// simluate work
			time.Sleep(500 * time.Millisecond)
			result <- n * n
		}
	}()

}

func main() {
	in := generator(1, 2, 3, 4, 5)

	numWorks := 3
	wg := sync.WaitGroup{}

	result := make(chan int)
	for i := 0; i < numWorks; i++ {
		wg.Add(1)
		go worker(in, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println(res)
	}
}
