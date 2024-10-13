package main

import (
	"fmt"
)

func main() {
	myArr := [5]int {}
	
	for i := 0; i < 5; i++ {
		myArr[i] = i
	}

	for i, v := range myArr {
		fmt.Printf("%v - %T - %v", v, v, i)
	}
}

/*
Hands-on exercise #42
● Using a COMPOSITE LITERAL:
● create an ARRAY which holds 5 VALUES of TYPE int
● assign VALUES to each index position.
● Range over the array and print the values out.
○ print out the VALUE and the TYPE
https://go.dev/play/p/aghCIHXlIsB
curriculum item # 106-hands-on-exercise-42
*/