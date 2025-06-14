package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// go routine, with time to make sure that routine executes before main terminates

	go func() {
		fmt.Println("Inside a go routine")
	}()

	time.Sleep(1 * time.Second)

	// using  wait groups
	// A WaitGroup is a synchronization primitive that lets you wait for a group of goroutines to finish.

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
	}

	wg.Add(len(words))

	for i, val := range words {

		go func() {
			defer wg.Done()
			fmt.Println(i, val)
		}()
	}

	wg.Wait()

	wg.Add(1)
	Demo(&wg)
	wg.Wait()

}
