package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// gen return receive channel
func gen() <-chan int {
	c := make(chan int)
	go func(){
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	// will block if c is not closed
	for v := range c {
		fmt.Println(v)
	}
}

/*
Hands-on exercise #3
● Starting with this code, pull the values off the channel using a for range loop
solution: https://play.golang.org/p/D3N4Tq54SN
*/