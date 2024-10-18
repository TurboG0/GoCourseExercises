package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	// only prints 42
	fmt.Println(<-c)
	// since there is 1 value left and it is 43 this will print 43
	fmt.Println(<-c)
}
