package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(x))
}


func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	// knows that it is returning total
	// recommended to not do returns like this
	return
}

// if scope is large longer variable names are recommended
// for more readibility
