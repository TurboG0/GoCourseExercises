package main

import "fmt"

func main() {
	// by making chanel like this, I'm only able to receive only
	// values onto the channel
	c := make(<-chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------")
	// chan int
	fmt.Printf("%T\n", c)
}
