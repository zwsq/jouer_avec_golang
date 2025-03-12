package main

import "fmt"

// Recursion is a programming concept where a function calls itself to solve a problem
// by breaking it down into smaller subproblems. In Go, recursion is commonly used for
// problems like factorial computation, Fibonacci sequence generation, and tree traversal.

// Recursive function to calculate factorial
func factorial(target int) int {
	if target == 0 {
		return 1
	}
	return target * factorial(target-1)
}

func main() {
	fmt.Println(factorial(5)) // Output: 120 (5*4*3*2*1)
}
