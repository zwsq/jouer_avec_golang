package main

import (
	"fmt"
	"sync"
)

func simpleRunner() {

	zooj := make(chan int)
	fard := make(chan int)
	fanIn := make(chan int)

	go send(zooj, fard)
	go receive(zooj, fard, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}
}
func send(z, f chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			z <- i
		} else {
			f <- i
		}
	}
	close(z)
	close(f)
}

func receive(z, f <-chan int, fan chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range f {
			fan <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range z {
			fan <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fan)
}
