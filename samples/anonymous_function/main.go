package main

import "fmt"

func main() {
	// Anonymous function signature:
	// func (p parameter(s)) (r return(s)){<code>}
	func(s string) {
		fmt.Println("This is my name from an anonymous function:", s)
	}("ZWSQ")
}
