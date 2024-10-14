package main

import (
	"fmt"
)

type person struct {
	first				string
	last				string
	favIcecreamFlavour	[]string
}

func main() {
	p1 := person {
		first:	"Saba",
		last:	"Dzodzuashvili",
		favIcecreamFlavour: []string{"Chocolate, banana, Nutela with nuts"},
	}
	p2 := person {
		first:	"Lui",
		last:	"Dzodzuashvili",
		favIcecreamFlavour: []string{"Vanilla, pashteti, jele"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIcecreamFlavour)
	fmt.Println(p2.favIcecreamFlavour)

	for _, v := range p1.favIcecreamFlavour {
		fmt.Println(p1.first, "Favourite icecream flavours are", v)
	}

	for _, v := range p2.favIcecreamFlavour {
		fmt.Println(p2.first, "Favourite icecream flavours are", v)
	}
}

/*
Hands-on exercise #53 - struct with slice
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
https://go.dev/play/p/nLcea3bIb7h
curriculum item # 126-hands-on-exercise-53-struct-with-slice
*/