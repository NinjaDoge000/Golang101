package main

import (
	"fmt"
	"time"
)

/*
	The done channel pattern is used to signal cancellation, timeout, or
	completion between goroutines. It's especially useful to avoid goroutine leaks by ensuring that
	all goroutines have a clean way to exit early when they're no longer needed.
*/

func doWork(done <-chan struct{}) {

	for {
		select {
		case <-done:
		default:
			fmt.Println("Doning Some Work...")
		}
	}

}

func main() {

	done := make(chan struct{})
	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done)

}
