package main

import (
	"fmt"
	"sync"
)

var msg string
var wg2 sync.WaitGroup

func printSomething(s string, mtx *sync.Mutex) {
	defer wg2.Done()
	mtx.Lock()
	msg = s
	mtx.Unlock()
	fmt.Println(msg)
}

func Demo(wg *sync.WaitGroup) {
	defer wg.Done()

	var mtx sync.Mutex

	wg2.Add(2)
	go printSomething("hello, world", &mtx)
	go printSomething("GoodBye, cruel world", &mtx)
	wg2.Wait()
	fmt.Printf("value of msg is %s", msg)

}
