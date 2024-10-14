package main

import "fmt"

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	x := doMath(3, 4, add)
	y := doMath(3, 4, sub)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("Printing out types of these functions")
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)
}