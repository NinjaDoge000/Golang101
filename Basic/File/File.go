package main

import (
	"fmt"
	"os"
)

const fileName = "file.txt"

func main() {

	// write a string to a file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("Hello, File!")
	fmt.Println("Data written to file successfully!")

	// read contents in file.txt
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read from file:", string(data))
}
