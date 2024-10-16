package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// addT is generic function that accepts int and float64 types
// T can be anything you want but it is common practice to name T
// func addT[T int | float64](a, b T) T{
// 	return a + b
// }

type mynumbers interface {
	//~int | ~float64

	//
	constraints.Integer | constraints.Float
}

// abstract previously ussed addT function with interface
func addT[T mynumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var n myAlias = 22
	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))
}
