package main

import "fmt"

// nested loops
func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("Iteration %v from loop I ,iteration %v from loop II\n", i, j)
		}
		fmt.Println()
	}
}
