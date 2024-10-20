package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

	//can do also like this
	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)

}

func foo() int {
	return 69
}

func bar() (int, string) {
	return 420, "Ha Ha Ha"
}

/*
Hands-on exercise #58 - basic funcs
● Hands on exercise
Todd McLeod - Learn To Code Go - Page 98
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
https://go.dev/play/p/Jeh_wripOEe
*/
