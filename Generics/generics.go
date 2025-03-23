package main

import "fmt"

// Constraint: T must be an integer
type Number interface {
	int | int64 | float64
}

// Generic function with constrained types
func Add[T Number](a, b T) T { // generics is like a template in C++
	return a + b
}

func main() {
	fmt.Println(Add(5, 10))
	fmt.Println(Add(3.5, 2.2))
	// fmt.Println(Add("a", "b")) // Compile error (string is not a Number)
}

//  func (receiver) methodName[T Type](args) returnType
