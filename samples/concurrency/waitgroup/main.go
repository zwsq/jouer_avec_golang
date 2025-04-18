package main

import (
	"fmt"
	"sync"
)

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("    Foo", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("  Bar", i)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
