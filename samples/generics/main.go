package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//--------------------------------------------------------------

// Generic function
func add1[T float64 | int](a, b T) T {
	return a + b
}

//--------------------------------------------------------------

// Type set
type number interface {
	float64 | int
}

// Generic fucntion with type set
func add2[T number](a, b T) T {
	return a + b
}

//--------------------------------------------------------------

// Alias
type myNumber int

// To accept aliases with certain underlying types, we add ~ before the accepted types
type validNumbers interface {
	~float64 | ~int
}

func add3[T validNumbers](a, b T) T {
	return a + b
}

//--------------------------------------------------------------

// Using constraints
type positiveNumbers interface {
	constraints.Float | constraints.Integer
}

func add4[T positiveNumbers](a, b T) T {
	return a + b
}

// --------------------------------------------------------------
func main() {
	var myNumber1 myNumber = 1395

	fmt.Println(add1(1, 3))
	fmt.Println(add2(1, 6))
	fmt.Println(add3(myNumber1, 6))
	fmt.Println(add4(230, 6))
}
