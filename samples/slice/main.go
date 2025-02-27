package main

import (
	"fmt"
)

func main() {
	exercise48()

}

// looping on a slice
func exercise42() {
	arrayOne := [5]int{}
	for i := range len(arrayOne) {
		arrayOne[i] = i
	}
	fmt.Println(arrayOne)
}

// use make to create a  slice
func exercise43() {
	slice43 := make([]int, 10)
	initNumber := 42
	for i := 0; i < 10; i++ {
		slice43[i] = initNumber + i
	}
	for v := range len(slice43) {
		fmt.Println(slice43[v])
	}
}

// using specific parts of a slice
func exercise44() {
	slice44 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//[inclusive:exclusive]
	fmt.Println(slice44[:5])
	fmt.Println(slice44[5:])
	fmt.Println(slice44[2:7])
	fmt.Println(slice44[1:6])

}

// appending to a slice
func exercise45() {
	slice45 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice45 = append(slice45, 52)
	slice45 = append(slice45, 53, 54, 55)
	slice451 := []int{56, 57, 58, 59, 60}
	slice45 = append(slice45, slice451...) //appending a slice to the end of another slice
	fmt.Println(slice45)
}

// using append to delete from a slice
func exercise46() {
	slice46 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice46 = append(slice46[:3], slice46[6:]...)
	fmt.Println(slice46)
}

// slice's length and capacity
func exercise47() {
	states := []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "California",
		"Colorado", "Connecticut", "Delaware", "Florida", "Georgia",
		"Hawaii", "Idaho", "Illinois", "Indiana", "Iowa",
		"Kansas", "Kentucky", "Louisiana", "Maine", "Maryland",
		"Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri",
		"Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio",
		"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont",
		"Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming",
	}

	for i := 0; i < cap(states); i++ {
		fmt.Printf("Value: \t%v index: \t%v\n", states[i], i)
	}
	fmt.Printf("The length is: %v and the cap is: %v\n", len(states), cap(states))

}

// multi-dimensional slice
func exercise48() {
	slice481 := []string{"james", "bond", "aktof"}
	slice482 := []string{"miss", "akbar", "hello"}
	multiDimensional := [][]string{slice481, slice482}
	fmt.Println(multiDimensional)
	for i, v := range multiDimensional {
		for ii, vv := range v {
			fmt.Printf("Index %v of Slice %v\t Value: %v\n", ii, i, vv)
		}
	}
}
