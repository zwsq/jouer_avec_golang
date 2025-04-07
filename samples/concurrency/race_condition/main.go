package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func unHandledRaceNoWg(limit int) {
	counter := 0
	start := time.Now()
	for i := 0; i < limit; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
		}()
	}
	elapsed := time.Since(start)
	fmt.Println("Unhandled without waitgroup:\t", counter, "\t", elapsed)
}

func unHandledRace(limit int) {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(limit)
	start := time.Now()
	for i := 0; i < limit; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Unhandled with waitgroup:\t", counter, "\t", elapsed)
}

func handledRaceWithMutex(limit int) {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(limit)
	start := time.Now()
	for i := 0; i < limit; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Handled with mutex:\t\t", counter, "\t", elapsed)
}

func handledRaceWithAtomic(limit int) {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(limit)
	start := time.Now()
	for i := 0; i < limit; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Handled with Atomic:\t\t", counter, "\t", elapsed)
}
func main() {
	fmt.Println()
	upperLimit := 5000000
	unHandledRaceNoWg(upperLimit)
	unHandledRace(upperLimit)
	handledRaceWithAtomic(upperLimit)
	handledRaceWithMutex(upperLimit)

}
