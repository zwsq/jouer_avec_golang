package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.IntN(5)
		switch x {
		case 0:
			fmt.Println("It's Zero.")
		case 1:
			fmt.Println("It's one.")
		case 2:
			fmt.Println("It's two.")
		case 3:
			fmt.Println("It's three.")
		case 4:
			fmt.Println("It's four.")
		default:
			fmt.Println("It's out of the range!")
		}
	}
}
