package main

import "fmt"

// looping in ranges (SLICES)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index %v: \t %v\n", i, v)
	}
}
