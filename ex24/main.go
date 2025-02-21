package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	switchy(20)
}
func switchy(iteration int) {
	for i := 1; i <= iteration; i++ {

		max := 10
		x := rand.IntN(max)
		y := rand.IntN(max)
		fmt.Printf("The value for x is %v for iteration %v\n", x, i)
		fmt.Printf("The value for y is %v for iteration %v\n", y, i)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both less than 4")
		case x > 6 && y > 6:
			fmt.Println("Both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("Both are between 4 and 6 inclusive")
		case y != 5:
			fmt.Println("Y is not equal to 5")
		default:
			fmt.Println("None of the conditions met!")
		}
		fmt.Println()
	}
}
