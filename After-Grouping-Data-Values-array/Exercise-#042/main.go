package main

import (
	"fmt"
)

func main() {
	myArr := [...]int {1, 2, 3, 4, 5}
	
	for _, v := range myArr {
		fmt.Printf("Value: %v\n", v)
	}
	fmt.Printf("Type of array is: %T", myArr)
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