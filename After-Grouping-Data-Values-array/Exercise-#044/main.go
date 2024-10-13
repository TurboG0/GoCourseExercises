package main

import (
	"fmt"
)

func main() {
	mySlice := []int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("1. ", mySlice[:5])
	fmt.Println("2. ", mySlice[5:])
	fmt.Println("3. ", mySlice[2:7])
	fmt.Println("4. ", mySlice[1:6])
}

/*
Hands-on exercise #44
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]
https://go.dev/play/p/LoAQoGHT-W6
curriculum item # 108-hands-on-exercise-44
*/