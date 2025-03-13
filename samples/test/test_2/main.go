package main

import "fmt"

func main() {
	a := 95
	b := 73
	fmt.Println(doMathematics(a, b, add))
	fmt.Println(doMathematics(a, b, subtract))
}

func add(a int, b int) (result int) {
	return a + b
}

func subtract(a int, b int) (result int) {
	return a - b
}

func doMathematics(a int, b int, f func(a int, b int) int) int {
	return f(a, b)
}
