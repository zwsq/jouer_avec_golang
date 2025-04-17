package main

import (
	"fmt"
	"log"
)

func main() {
	zooj := make(chan int)
	fard := make(chan int)
	quit := make(chan int)

	log.Print("Saving values to the channels")
	go send(zooj, fard, quit)
	log.Print("Saved values to the channels")
	log.Print("Receiving values from the channels")
	receive(zooj, fard, quit)
	log.Print("Received values from the channels")

}

func send(z, f, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			z <- i
		} else {
			f <- i
		}
	}
	q <- 1020
}

func receive(z, f, q <-chan int) {
	for {
		select {
		case v := <-z:
			fmt.Println("Zooj", v)
		case v := <-f:
			fmt.Println("Fard", v)
		case v := <-q:
			fmt.Println("Quit", v)
			return
		}
	}

}
