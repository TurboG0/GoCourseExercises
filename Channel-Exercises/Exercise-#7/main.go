package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int)
	const goroutines = 10

	for i := 0; i < goroutines; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- rand.Intn(10)
			}
		}()
	}
	for i := 0; i < goroutines * 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("About to exit")
}

/*
Hands-on exercise #7
● write a program that
○ launches 10 goroutines
■ each goroutine adds 10 numbers to a channel
○ pull the numbers off the channel and print them
*/
