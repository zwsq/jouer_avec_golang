package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

func main() {
	// /unbufferedChannel()
	bufferedChannel(4)
}

func unbufferedChannel() {
	ch := make(chan int)
	go func() { ch <- 32 }()
	fmt.Println(<-ch)

}

func bufferedChannel(volume int) {
	ch := make(chan int, volume)
	for i := 0; i < volume; i++ {
		go func() {
			ch <- rand.Int()
			fmt.Println("Saved value:", i+1)
		}()
	}
	log.Print("Values saved to the buffered channel.")

	for i := 0; i < volume; i++ {
		fmt.Println(<-ch)
	}
}
