package main

import "fmt"

func main() {
	// function expression is when we assign a function to a variable
	x := func() {
		fmt.Println("X")
	}

	y := func(s string) {
		fmt.Println("This is y:", s)
	}

	x()
	y("DJ")
}
