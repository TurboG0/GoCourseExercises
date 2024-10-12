package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("Value of x is: %d\nValue of y is: %d\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("Both x and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both x and y are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("X is in between of 4 and 6")
	} else if y != 5 {
		fmt.Println("y is not equal to 5")
	} else {
		fmt.Println("non of the previous cases meet")
	}
}