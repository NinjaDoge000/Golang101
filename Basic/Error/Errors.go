package main

import (
	"errors"
	"fmt"
)

func divide(a float64, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Cannot divide by Zero")
	}

	return a / b, nil

}

func main() {

	/*

		1. Errors are handled explicitly. Functions that can fail return an error as a second return value
		   Forces the caller to handle the error.

		2. Errors are values. They are just another value that a function can return.

		3. Unlike languages like Java, Python, or C++, Go does not have exceptions.

		4. Go uses the built-in error type, which is an interface. Custom error types can be created by implementing the Error() method.

		5. Panic and Recover are used for exceptional cases, not for regular error handling. Read more about it!

	*/

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("the Result is ", result)

}
