package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Dildo",
		Last:  "Beggins",
		Age:   69,
	}

	p2 := person{
		First: "Saba",
		Last:  "Dzodzuashvili",
		Age:   22,
	}

	people := []person{
		p1,
		p2,
	}

	bs, err := json.Marshal(people)

	if err != nil {
		log.Fatalf("error - %v\n", err)
	}

	fmt.Println(string(bs))
}
