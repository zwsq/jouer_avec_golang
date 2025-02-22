package main

import "fmt"

// Looping in ranges (Maps)
func main() {
	myMap := map[string]int{
		"birthday":     19950121,
		"nationalCode": 17791510,
	}

	for k, v := range myMap {
		fmt.Printf("%v\t %v\n", k, v)
	}
}
