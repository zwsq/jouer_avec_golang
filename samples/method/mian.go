package main

import "fmt"

type person struct {
	name string
}

// A method in Go is a function that has a receiver.
// The receiver is a special parameter that allows
// a function to be associated with a specific struct (or type),
// enabling object-oriented-like behavior
func (p person) sayYourName() {
	fmt.Println("My name is", p.name)
}

func main() {
	zwsq := person{
		name: "Abolfazl",
	}
	zwsq.sayYourName()
}
