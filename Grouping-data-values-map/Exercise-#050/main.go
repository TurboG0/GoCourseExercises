package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string][]string)
	myMap[`bond_james`] = []string {`shaken, not stirred`, `martinis`, `fast cars`}
	myMap[`moneypenny_jenny`] = []string {`intelligence`, `literature`, `computer science`}
	myMap[`no_dr`] = []string {`cat`, `ice cream`, `sunset`}
	fmt.Printf("%#v\n", myMap)
	
	for k, v := range myMap {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

	fmt.Println("------------------------")
	myMap[`fleming_ian`] = []string {`steaks`, `cigars`, `espionage`}
	for k, v := range myMap {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}

/*
Hands-on exercise #50 - add a record
Using the code from the previous example, add a record to your map. Now print the map out
using the “range” loop
key value
`fleming_ian` `steaks`, `cigars`, `espionage`
https://go.dev/play/p/AJ8ZsHCC0wU
curriculum item # 119-hands-on-exercise-50
*/