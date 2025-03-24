package main

import "fmt"

// Generic function
func add1[T float64 | int](a, b T) T {
	return a + b
}

// Type set
type number interface {
	float64 | int
}

// Generic fucntion with type set
func add2[T number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add1(1, 3.3))
	fmt.Println(add2(1, 6.3))
}
