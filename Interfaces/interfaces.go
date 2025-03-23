package main

import (
	"fmt"
)

// Define the Speaker interface
type Speaker interface {
	Speak() string
}

type Animal interface {
	Speaker // Embed the Speaker interface
	Move() string
}

// Implement Speak for Dog
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s is barking", d.Name)
}

func (d Dog) Move() string {
	return fmt.Sprintf("%s is running", d.Name)
}

// Implement Speak for Cat
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s is meowing", c.Name)
}

// Function that works with any Speaker
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	var d = Dog{Name: "Tommy"}
	MakeSpeak(d)
	var c = Cat{Name: "Kitty"}
	MakeSpeak(c)

	// Using the Animal interface to call Speak and Move and Dog's Speak and Move methods
	// This is possible because Dog implements the Animal interface
	var a Animal = Dog{Name: "Tommy"}
	fmt.Println(a.Speak())
	fmt.Print(a.Move())
}
