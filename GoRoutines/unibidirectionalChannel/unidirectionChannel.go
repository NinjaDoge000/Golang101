package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// read-only channel
func receiveData(ch <-chan int) {

	defer wg.Done()

	for data := range ch {
		fmt.Println("value of data is", data)
	}

	_, ok := <-ch

	if !ok {
		fmt.Println("channel is closed")
	}
}

// write-only channel
func sendData(ch chan<- int) {

	defer close(ch)
	defer wg.Done()

	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func unidirectionalChannel() {

	wg.Add(2)

	ch := make(chan int)

	go receiveData(ch)
	go sendData(ch)

	wg.Wait()

}
