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

	age := myMap["nationalCode"]
	fmt.Println("age 1", age)

	age = myMap["Q"]
	fmt.Println("age2", age)

	// comma ok example
	if v, ok := myMap["Q"]; !ok {
		fmt.Println("No Q found", v)
	}

}
