package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64 = 0

	const gs = 100
	var wg sync.WaitGroup

	//	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", atomic.LoadInt64(&counter))
}

/*
Hands-on exercise #5
Fix the race condition you created in exercise #4 by using package atomic
code: https://github.com/GoesToEleven/go-programming
*/
