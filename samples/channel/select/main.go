package main

import (
	"fmt"
	"log"
)

func main() {
	zooj := make(chan int)
	fard := make(chan int)
	quit := make(chan bool)

	log.Print("Saving values to the channels")
	go send(zooj, fard, quit)
	log.Print("Saved values to the channels")
	log.Print("Receiving values from the channels")
	receive(zooj, fard, quit)
	log.Print("Received values from the channels")

}

func send(z, f chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			z <- i
		} else {
			f <- i
		}
	}
	close(q)
}

func receive(z, f <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-z:
			fmt.Println("Zooj", v)
		case v := <-f:
			fmt.Println("Fard", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("Exiting", i)
				return
			} else {
				fmt.Println("I should not be here", i)
			}
		}
	}

}
