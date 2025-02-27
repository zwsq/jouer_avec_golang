package main

import (
	"fmt"

	mol "github.com/zwsq/g"
	"golang.org/x/exp/rand"
)

func main() {
	// If else statement
	name := "Fatima"
	if name == "Fatima" {
		fmt.Println(mol.Capitalizer(name))
	} else if name == "ZWSQ" {
		fmt.Println(mol.Multiplier(2, 3))
	} else {
		fmt.Println("Not met")
	}

	// Switch case
	switch name {
	case "32":
		fmt.Println("32")
	case "Fatima":
		fmt.Println("HEEEEEEEEEEEEEEEEEEEE")
		fallthrough
	default:
		fmt.Println("None of the cases met")
	}

	// Statement statement idiom
	maximum_value := 50
	if x := 40 * rand.Intn(maximum_value); x <= 32 {
		fmt.Println("Statement statement idiom condition met")
	} else {
		fmt.Println(x)
	}

	ifif := rand.Intn(maximum_value)
	fmt.Printf("The name of the file is ifif and the value is %v\n", ifif)
	if ifif <= 100 {
		fmt.Println("Between 0 and 100")
	} else if 100 < ifif && ifif <= 200 {
		fmt.Println("Between 100 and 200")
	} else {
		fmt.Println("Between 200 and 250")
	}

	fmt.Println(ifif)
	fmt.Println(rand.Intn(maximum_value))
	fmt.Println(rand.Intn(maximum_value))
	fmt.Println(rand.Intn(maximum_value))
	fmt.Println(rand.Intn(maximum_value))

	nu := []int{3, 3, 2}
	twoSum(nu, 5)
}

func twoSum(nums []int, target int) []int {
	answer := []int{}
	for i := range nums {
		if nums[i]+nums[i+1] == target {

		}
	}
	return answer

}
