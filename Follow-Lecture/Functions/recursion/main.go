package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}

//anything you can do with recursion you can do also with loops
//loops are easier to understand and more performant