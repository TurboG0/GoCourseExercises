package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin Goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hi, I'm Goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hi, I'm Goroutine 2")
		wg.Done()
	}()

	fmt.Println("Mid CPUs:", runtime.NumCPU())
	fmt.Println("Mid Goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPUs:", runtime.NumCPU())
	fmt.Println("End Goroutines", runtime.NumGoroutine())

}

/*
Hands-on exercise #1
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
code: https://github.com/GoesToEleven/go-programming
*/
