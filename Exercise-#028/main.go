package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("Value of x is: %d\nValue of y is: %d\n", x, y)
	
	switch {
	case x < 4 && y < 4:
		fmt.Println("Both x and y are less than 4")
	case x > 6 && y > 6:
		fmt.Println("Both x and y are greater than 6")
	case y != 5:
		fmt.Println("y is not equal to 5")
	default:
		fmt.Println("non of the previous cases meet")
	}
}