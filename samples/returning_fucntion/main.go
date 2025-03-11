package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)
	b := bar()
	fmt.Printf("Before calling: %T and after calling %T\n", b, b())
}

func foo() int {
	return 73
}

func bar() func() int {
	return func() int {
		return 73
	}
}
