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
	
}

/*
Hands-on exercises
Hands-on exercise #49 - map[string][]string
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
key value
`bond_james` `shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` `intelligence`, `literature`, `computer science`
`no_dr` `cats`, `ice cream`, `sunsets`
*/

/*
{`shaken, not stirred`, `martinis`, `fast cars`}
{`intelligence`, `literature`, `computer science`}
{`cat`, `ice cream`, `sunset`}
*/