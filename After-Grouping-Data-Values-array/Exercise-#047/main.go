package main

import (
	"fmt"
)

func main() {
	stateNames := make([]string, 0, 50)
	fmt.Println("Length: ", len(stateNames))
	fmt.Println("Capacity: ", cap(stateNames))

	stateNames = append(stateNames, 
		`Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, 
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, 
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, 
		` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, 
		` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	)
	fmt.Println("--------------------------------")
	fmt.Println("after adding elements:")
	fmt.Println("Length: ", len(stateNames))
	fmt.Println("Capacity: ", cap(stateNames))
	for i := 0; i < len(stateNames); i++ {
		fmt.Printf("Index: %v\tValue: %v\n", i, stateNames[i])
	}
}

/*
Hands-on exercise #47
For this exercise, do the following:
● Create a slice to store the names of all of the states in the United States of America.
○ Use make and append to do this.
○ Goal: do not have the array that underlies the slice created more than once.
● Print out
○ the len
○ the cap
○ the values, along with their index position, without using the range clause.
● Here is a list of the 50 states:
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
https://go.dev/play/p/-cwlxRpnJAw
curriculum item # 111-hands-on-exercise-47
*/