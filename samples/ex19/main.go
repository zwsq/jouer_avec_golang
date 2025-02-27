package main

import (
	"fmt"

	"math/rand/v2"
)

func main() {
	maximum_value := 250
	val := rand.IntN(maximum_value)
	fmt.Printf("The name of the file is val and the value is %v\n", val)
	if val <= 100 {
		fmt.Println("Between 0 and 100")
	} else if 100 < val && val <= 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}
}
