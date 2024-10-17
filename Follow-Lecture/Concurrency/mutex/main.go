package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	// adds gs amount of wait groups
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}
