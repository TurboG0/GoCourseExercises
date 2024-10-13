package main

import (
	"fmt"
)

func main() {
	
	firstSlice := []string{"James", "Bond", "Shaken, not stirred"}
	secondSlice := []string{"Miss", "Moneypenny", "I'm 008."}

	slice2D := [][]string{firstSlice, secondSlice}

	for i, v := range slice2D {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}

/*
Hands-on exercise #48
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
https://go.dev/play/p/2jxGMyXiZrE
curriculum item # 112-hands-on-exercise-48
*/