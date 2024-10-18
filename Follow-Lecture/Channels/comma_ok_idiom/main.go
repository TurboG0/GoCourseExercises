package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("About to exit")
}

func send(eve, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(eve, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-eve:
			fmt.Println("the value is received from the even channel", v)
		case v := <-odd:
			fmt.Println("the value is received from the odd channel", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}
