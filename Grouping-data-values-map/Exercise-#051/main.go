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

	fmt.Println("------------------------")
	delete(myMap, `bond_james`)
	for k, v := range myMap {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}

/*
Hands-on exercise #51 - delete a record
Using the code from the previous example, delete a record from your map. Now print the map
out using the “range” loop
https://go.dev/play/p/yba3zekRiBY
curriculum item # 120-hands-on-exercise-51
*/