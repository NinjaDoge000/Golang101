package main

import (
	"fmt"
)

// Define the Speaker interface
type Speaker interface {
	Speak() string
}

// Implement Speak for Dog
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s is barking", d.Name)
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
}
