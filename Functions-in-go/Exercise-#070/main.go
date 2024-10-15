package main

import "fmt"

func main() {
	f := outer()
	fmt.Println(f())
}

func outer() func() int {
	return func() int {
		return 42
	}
}

/*
Hands-on exercise #70 - func return
● Create a func
○ which returns a func
■ which returns 42
● assign the returned func to a variable
● call the returned func
● print
https://go.dev/play/p/eqpKJ5jCv3g
*/