// Info
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go foo()
	go bar()
	fmt.Println(runtime.NumCPU())
	wg.Wait()
}

func foo() {
	for i := 1; i < 11; i++ {
		fmt.Println("Foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i < 11; i++ {
		fmt.Println("Bar ", i)
	}
	wg.Done()
}
