package main

import "fmt"

func main() {
	c := make(chan int, 1)
	
	c <- 42
	// we only have room for 1 so c <- 43 will cause error
	c <- 43
	
	fmt.Println(<-c)

}