package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop in go is flexible and can mimic while loop

	// for loop with only condition (like while loop) and its called conditional loop
	j := 0
	for j < 5 {
		fmt.Println("The value of j is", j)
		j++
	}

	// infinite loop
	for {
		fmt.Println("This is an infinite loop")
		break
	}

	// switch statement
	// notice that the break statement is not required and not allowed
	choice := 3
	switch choice {
	case 1:
		fmt.Println("You chose 1")
	case 2:
		fmt.Println("You chose 2")
	case 3, 4:
		fmt.Println("You chose 3 or 4")
	default:
		fmt.Println("You chose something else")
	}
}
