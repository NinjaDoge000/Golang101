package main

import (
	"fmt"

	mathutils "Golang101/Math" // import the package with an alias

	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Absolute value of -10 is", mathutils.Abs(-10))
	fmt.Println("Factorial of 5 is", mathutils.Factorial(5))
	// fmt.Println("2^3 is", mathutils.Pow(2, 3)) -> this will not work as Pow is not exported

	// using a third party package
	fmt.Println("Silly name:", randomdata.SillyName())
}
