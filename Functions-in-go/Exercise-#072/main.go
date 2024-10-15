package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
/*
Hands-on exercise #72 - closure
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:
https://go.dev/play/p/tEjP1-6S2LW
*/