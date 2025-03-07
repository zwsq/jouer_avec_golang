package main

import "fmt"

func main() {

	// The deferred function call is not executed until the surrounding function returns or panics
	defer variadic(1, 2, 3, 4, 56, 23, 4, 23)

	foo()

	bar("ZWSQ")

	whoAmI := aloha("ZWSQ")
	fmt.Println(whoAmI)

	message, birthYear := hopsh("ZWSQ", 30)
	fmt.Println(message, birthYear)

	sliceOfInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// This is called unfurling a slice
	variadic(sliceOfInt...)

}

// no parameter, no return
func foo() {
	fmt.Println("No params, no returns")
}

// one parameter, no return
func bar(s string) {
	fmt.Println("Hello dear ", s)
}

// one parameter, one return
func aloha(s string) string {
	return fmt.Sprint("You are ", s, " as far as I know!")
}

// two params, two returns
func hopsh(name string, age int) (message string, birthYear int) {
	bd := 2025 - age
	return fmt.Sprint("Thank you for using me ", name), bd
}

// A variadic function takes N number of type T
func variadic(x ...int) (sum int) {
	sum = 0
	for _, v := range x {
		sum += v
	}
	println("The sum is", sum)
	return sum
}
