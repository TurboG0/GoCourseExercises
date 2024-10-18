package main

import "fmt"

func main() {
	// buffer channel
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------")
	// chan int
	fmt.Printf("%T\n", c)
}