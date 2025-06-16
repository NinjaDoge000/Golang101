package main

import "fmt"

func main() {

	// Zero values
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	// print the values of a, b, c, d, and e
	fmt.Println("the value of a is", a)
	fmt.Println("the value of b is", b)
	fmt.Println("the value of c is", c)
	fmt.Println("the value of d is", d)
	fmt.Println("the value of e is", e)

	// output:
	// the value of a is 0
	// the value of b is 0
	// the value of c is false
	// the value of d is
	// the value of e is <nil>
}
