package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("With function literal", <-c)

	//with buffer channel
	cb := make(chan int, 1)
	cb <- 42
	fmt.Println("With channel buffer", <-cb)
}

/*
Hands-on exercise #1
● get this code working:
○ using func literal, aka, anonymous self-executing func
■ solution: https://play.golang.org/p/SHr3lpX4so
○ a buffered channel
■ solution: https://play.golang.org/p/Y0Hx6IZc3U
*/
