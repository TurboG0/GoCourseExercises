package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)	
		
		switch x {
		case 0:
			fmt.Printf("Iteration: %v \t X is equal to 0\n", i)
		case 1:
			fmt.Printf("Iteration: %v \t X is equal to 1\n", i)
		case 2:
			fmt.Printf("Iteration: %v \t X is equal to 2\n", i)
		case 3:
			fmt.Printf("Iteration: %v \t X is equal to 3\n", i)
		case 4:
			fmt.Printf("Iteration: %v \t X is equal to 4\n", i)
		}
	}
}