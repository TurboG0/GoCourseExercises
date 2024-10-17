package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 2, 3, 1, 6, 9, 12}
	xs := []string{"Saba", "Gela", "Murtazi", "Skywanker", "Beggins"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println("Sorted", xi)

	fmt.Println("-----------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println("Sorted", xs)

}
