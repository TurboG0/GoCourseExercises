package main

import (
	"fmt"
)

func main() {
	mySlice := []int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range mySlice {
		fmt.Printf("%v \t %T \t %v\n", v, v, i)
	}
}

/*
Hands-on exercise #43
● Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign these 10 VALUES
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
● Range over the slice and print the values out.
○ print out the VALUE and the TYPE
https://go.dev/play/p/84r6QDOaMrs
curriculum item # 107-hands-on-exercise-43
*/