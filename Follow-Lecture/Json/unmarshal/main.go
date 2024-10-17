package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[
			{"First":"Dildo","Last":"Beggins","Age":69},
			{"First":"Saba","Last":"Dzodzuashvili","Age":22}
		]`

	var people []person
	bs := []byte(s)
	err := json.Unmarshal(bs, &people)

	fmt.Printf("s: %T\n", s)
	fmt.Printf("bs: %T\n", bs)

	if err != nil {
		log.Fatalf("error - %v\n", err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("\nPerson number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
