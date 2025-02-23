package main

import (
	"fmt"
)

func main() {

	// Arrays
	{
		as := [5]int{3, 4, 3, 5}
		fmt.Println(as)
	}

	//Slices
	{
		slice1 := []int{2, 3, 4, 5, 6}
		fmt.Println(slice1)
		slice2 := []int{}
		for v := range slice1 {
			slice2 = append(slice2, v*2)
		}
		fmt.Println(slice2)
	}
}
