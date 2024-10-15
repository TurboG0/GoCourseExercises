package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(x...)
	fmt.Println(sum)

	sum1 := bar(x)
	fmt.Println(sum1)
}

func foo(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

/*
Hands-on exercise #59 - variadic func
● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in
https://go.dev/play/p/AVSkvMuaz0H
*/
