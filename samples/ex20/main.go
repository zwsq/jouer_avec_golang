package main

import (
	"fmt"

	"math/rand/v2"
)

func main() {

	maximum_value := 250
	val := rand.IntN(maximum_value)
	fmt.Printf("The name of the file is val and the value is %v\n", val)

	switch {
	case val >= 0 && val <= 100:
		fmt.Println("Between 0 and 100")
	case val > 100 && val <= 200:
		fmt.Println("Between 101 and 200")
	case val > 200:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("Invalid")

	}
}
