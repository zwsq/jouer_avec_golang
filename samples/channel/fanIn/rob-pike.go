package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func robRunner() {
	c := fanIn(greeting("HBK"), greeting("Mamali"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Both of you are wasting my time, Goodbye.")
}

func greeting(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.IntN(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(i1, i2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-i1
		}
	}()
	go func() {
		for {
			c <- <-i2
		}
	}()
	return c
}
