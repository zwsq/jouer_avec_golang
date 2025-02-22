package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// withoutStatementStatement(100)
	withStatementStatement(1000000, 0, 1000)

}

func withoutStatementStatement(timesToRun int) {
	for i := 0; i < timesToRun; i++ {
		x := rand.IntN(5)
		if x == 3 {
			fmt.Println("The x value is three.")
		}

	}
}

// statement statement idiom
func withStatementStatement(timesToRun int, valueToFind int, maximumRandomValue int) {
	timesFound := 1
	for i := 0; i < timesToRun; i++ {
		if x := rand.IntN(maximumRandomValue); x == valueToFind {
			fmt.Printf("Value %v found in iteration number %v, Totally found %v\n", valueToFind, i, timesFound)
			timesFound++
		}
	}
}
