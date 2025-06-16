package main

import "fmt"

func main() {

	// variables
	var a int = 10         // explicit type declaration
	var b = 20             // type inference
	var c, d int = 30, 40  // multiple variable declaration
	var e, f = 50, "hello" // multiple variables of different types
	g := 60                // simple variable declaration
	g, h := 70, 80         //
	var r float64 = 2.0    // float variable

	// print the values of a and b
	fmt.Println("the value of a is", a, "and the value of b is", b)

	// format specifiers
	fmt.Printf("the value of a is %d and the value of b is %d\n", a, b)

	fmt.Println("the value of c is", c, "and the value of d is", d)

	fmt.Println("the value of e is", e, "and the value of f is", f)

	fmt.Println("the value of g is", g)

	fmt.Println("the value of g is", g, "and the value of h is", h)

	fmt.Println("the value of r is", r)

	// constants
	const pi = 3.14               // untyped constant
	const Euler float64 = 2.71828 // typed constant
	const (                       // multiple constant declaration
		x = 1
		y = 2
	)

}
