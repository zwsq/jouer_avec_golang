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

	if x < 4 && y < 4 {
		fmt.Println("Both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("X is 4, 5 or 6")
	} else if y != 5 {
		fmt.Println("Y is not equal to 5")
	} else {
		fmt.Println("None of the conditions met!")
	}

}
