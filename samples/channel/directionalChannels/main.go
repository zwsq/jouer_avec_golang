package main

import "fmt"

func main() {
	/*
		bc := make(chan int)

		// receive only channel: Only receiving values from this channel is allowed.
		roc := make(<-chan int)

		// send only channel: Only sending values to this channel is allowed.
		soc := make(chan<- int)

		// following code causes error because I'm trying to send a value to a receive only channel.
		roc <- 2

		// following code causes error because I'm trying to reveive a value from a send only channel.
		fmt.Println(<-soc)

		// Assigning a more specific type to a general type is not possible and would cause error like below
		bc = roc
		bc = soc

		// Assigning a more general type to a more specific type is possible like below
		roc = bc
		soc = bc
	*/
	c := make(chan int)

	go foo(c)
	bar(c)

}

// send value to the channel
func foo(c chan<- int) {
	c <- 73
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
