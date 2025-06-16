package main

import "fmt"

// function with parameters and return type
func addNumbers(a int, b int) int {
	return a + b
}

// funtion with parameter and no return type
func printResult(result int) {
	fmt.Println("The result is", result)
}

// function with no parameters and no return type
func printMessage() {
	fmt.Println("Goodbye!")
}

// function with multiple return values
func addSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

// variadic function, that takes a variable number of arguments
func printResults(results ...int) {
	fmt.Print("The results are")
	for _, result := range results {
		fmt.Printf(", %d", result)
	}
	fmt.Println()
}

// named return values
func addSubtractNamed(a int, b int) (sum int, difference int) {

	// no need to declare sum and difference as they are already declared in the function signature
	sum = a + b
	difference = a - b
	return sum, difference // or just return; the named return values are automatically returned
}

func main() {

	var c int = addNumbers(10, 20)
	var d, e int = addSubtract(10, 20)

	printResult(c)
	printResults(d, e)
	printMessage()

	// anonymous functions
	func() {
		fmt.Println("Hello!")
	}()

	// anonymous funtion assigned to a variable
	f := func() {
		fmt.Println("Hello!")
	}

	f()

}
