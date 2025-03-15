package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(powerUp(50, 5))
}

func powerUp(input int, power int) int {
	if power < 0 {
		log.Fatalln("power is invalid ")
	}
	output := 1
	for power != 0 {
		output *= input
		power--
	}
	return output
}
