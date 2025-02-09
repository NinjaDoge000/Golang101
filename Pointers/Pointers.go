package main

import "fmt"

func modifyValue(a *int) {
	*a = 20
	// next := a + 1 // you can't do this in go because you can't do pointer arithmetic
}

func main() {

	var a int = 10

	var b *int = &a
	var c **int = &b // pointer to a pointer

	// e := &a // shorthand for creating a pointer to a variable
	// d := &b // shorthand for creating a pointer to a pointer

	fmt.Println("Value of a is", a)
	fmt.Println("Address of a is", &a)
	fmt.Println("Address stored by b is", b)
	fmt.Println("Value stored by b is", *b)
	fmt.Println("Value stored by c is", **c)

	modifyValue(&a)
	fmt.Println("Value of a after modification is", a)

	var p *int // null value of a pointer is nil
	if p == nil {
		fmt.Println("p is nil")
	}
}
