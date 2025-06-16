package main

import (
	"fmt"
	"sync"
)

// make it unidrectional write only so that func cannot read
func addToChannel1(ch1 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch1)

	ch1 <- 42
}

func addToChannel2(ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch2)

	ch2 <- 42
}

func main() {

	var ch1, ch2 chan int = make(chan int, 1), make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go addToChannel1(ch1, &wg)
	go addToChannel2(ch2, &wg)

	wg.Wait()

	// select waits until one channel is ready.
	// select with no defualt means it will wait infinitely until atleast one is ready
	// if both are not ready defualt will hit
	// if both are ready, one will be radomly selected
	select {
	case <-ch1:
		fmt.Println("added to channel 1")
	case <-ch2:
		fmt.Println("added to channel 2")
	default:
		fmt.Print("nothing is ready!") // this will never hit because we wait for all routines to complete
	}

}
