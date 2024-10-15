package main

import "fmt"

func main() {
	// will be last
	defer fooDefer1()
	// will be 2nd last
	defer fooDefer2()
	// will be first
	foo()
}

func foo() {
	fmt.Println("Function without defer")
}

func fooDefer1() {
	fmt.Println("First function with defer")
}

func fooDefer2() {
	fmt.Println("Second function with defer")
}

// defer functions run into LIFO order
// last in first out

/*
Hands-on exercise #60 - defer func
● “defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run
https://go.dev/play/p/tbvSX8qy6TT
*/
