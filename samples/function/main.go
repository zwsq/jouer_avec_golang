package main

import "fmt"

func main() {

	foo()

	bar("ZWSQ")

	whoAmI := aloha("ZWSQ")
	fmt.Println(whoAmI)

	message, birthYear := hopsh("ZWSQ", 30)
	fmt.Println(message, birthYear)

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
