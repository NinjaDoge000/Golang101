package main

import "fmt"

func main() {

	// arrays in go

	var a [5]int // array of 5 integers
	a[0] = 1
	a[1] = 2

	fmt.Println(a) // -> [1 2 0 0 0]

	// shorthand for creating an array
	b := [3]int{1, 2, 3}
	fmt.Println(b) // -> [1 2 3]

	// slicing an array
	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c) // -> [1 2 3 4 5]
	d := c[1:3]    // -> [2 3] index 1, 2
	e := c[2:]     // -> [3 4 5]
	f := c[:3]     // -> [1 2 3] index: 0, 1, 2
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	shalloCopy := c
	fmt.Println("shalloCopy", shalloCopy)
	c[0] = 100
	fmt.Println("shalloCopy", shalloCopy)

	// dynamic array -> slice
	slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	var slice2 []float64

	fmt.Println(slice2)
	fmt.Println(slice[0])
	slice = append(slice, 6.6)
	fmt.Println(slice)
	slice = append(slice, 7.7, 8.8, 9.9)

	slice2 = append(slice2, slice...) // ... is used to append all elements of a slice to another slice

	// slice vs array
	// array is fixed size, slice is dynamic
	// array is passed by value, slice is passed by reference

}
