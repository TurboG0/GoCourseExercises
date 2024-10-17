package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())

	var counter int64 = 0
	const gs = 100
	var wg sync.WaitGroup
	// adds gs amount of wait groups
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//takes the value address that i want to change
			// and delta which means how much or what am i changing
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}
