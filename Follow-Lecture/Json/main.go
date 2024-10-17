package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// type Person struct {
// 	First string
// 	Last  string
// 	Age   int
// }

// from json-to-go
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// p1 := Person{
	// 	First: "James",
	// 	Last:  "Bond",
	// 	Age:   32,
	// }

	// p2 := Person{
	// 	First: "Saba",
	// 	Last:  "Dzodzuashvili",
	// 	Age:   22,
	// }

	// people := []Person{
	// 	p1,
	// 	p2,
	// }

	//fmt.Println(people)

	// if i want to send this data to someone i need to turn this to json
	// so i need to marshal it
	// bs, err := json.Marshal(people)

	// if err != nil {
	// 	log.Fatalf("Error - %v", err)
	// }

	// since json.Marshal returns []byte we need to convert it to string
	// or we will get numbers
	// and also in order for it to work struct should be exportable with
	// it's fields meaning struct identifier should start with capital letter
	//fmt.Println(string(bs))

	fmt.Println("---------------------------")
	
	// to use this i need to convert it into []bytes
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Saba","Last":"Dzodzuashvili","Age":22}]`
	bs := []byte(s)
	
	// people := []person {}
	var people []person

	err := json.Unmarshal(bs, &people)


	if err != nil {
		log.Fatalf("error - %v\n", err)
	}
	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
