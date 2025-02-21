package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	max := 10
	x := rand.IntN(max)
	y := rand.IntN(max)
	fmt.Printf("The value for x is %v of type %T\n", x, x)
	fmt.Printf("The value for y is %v of type %T\n", y, y)

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
}
