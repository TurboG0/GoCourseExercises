package main

import "fmt"

func main() {
	// why doesn't this run?
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)

}