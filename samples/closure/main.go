package main

import "fmt"

// In Go, closures are functions that capture and remember the variables from their surrounding scope, even after the scope has ended.
// Closures are a powerful feature in functional programming, allowing functions to retain access to variables defined outside of their body.
func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
