package main

import (
	"fmt"
	"sync"
)

// print go routine statments i order of execution

func orderedExec() {

	var wg sync.WaitGroup

	ch := make(chan struct{})

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ch <- struct{}{}
			fmt.Printf("printing the function %d \n", i)
		}()

		<-ch
	}

	wg.Wait()

}
