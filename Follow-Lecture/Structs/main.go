package main 

import (
	"fmt"
)

type person struct {
	first	string
	last	string
	age		int
}

//embedded struct
type secretAgent struct {
	person
	first 	string
	ltk 	bool
}

type foo int

func main() {
	//embedded struct
	sa1 := secretAgent{ 
		person: person{
			first: "Saba",
			last:  "Dzodzu",
			age: 22,
		},
		first: "Gela",
		ltk: true,
	}
	p2 := person{ 
		first: "Dildo",
		last: "Beggins",
		age: 420,
	}
	anonStruct := struct {
		name		string
		lastName	string
		age			int
		p_ID		string
	}{
		name:		"Turbo",
		lastName: 	"Gelashvili",
		age:		96,
		p_ID:		"99999988865",
	}

	p3 := person{ 
		first: "Luke",
		last: "Skywanker",
		age: 69,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Printf("%T \t %#v\n",p3, p3)

	fmt.Println("--------------------------")
	fmt.Println(sa1)
	fmt.Println(sa1.person.first, sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person)
	fmt.Printf("%T\t%#v\n", sa1, sa1)
	
	fmt.Println("--------------------------")
	fmt.Printf("Anon struct: %T \t %#v\n", anonStruct, anonStruct)
	fmt.Println("--------------------------")
}

// you don't create classes in go, you create types
// why would someone create types like this: type foo int?
// you do it to make code to be self documented and more readable
// and attach some methods to it.

// There is padding and architectural alignment for structs
// if you want to priritize performanze lay the fields out from
// largest to smallest, eg int64, float32, bool