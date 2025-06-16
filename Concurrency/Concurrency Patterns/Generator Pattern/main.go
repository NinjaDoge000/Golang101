package main

import (
	"fmt"
	"math/rand"
)

/*
	A generator is a function that starts a goroutine,
	pushes values onto a channel, and returns that channel to the caller.

*/

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {

	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			default:
				stream <- fn()
			}
		}
	}()

	return stream
}

func main() {

	done := make(chan bool)
	randomNumFetcher := func() int { return rand.Intn(50000000) }

	for rando := range repeatFunc(done, randomNumFetcher) {
		fmt.Println(rando)
	}
}
