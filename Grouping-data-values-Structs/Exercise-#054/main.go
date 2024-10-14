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

	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Println(p1.favIcecreamFlavour)
	// fmt.Println(p2.favIcecreamFlavour)

	// for _, v := range p1.favIcecreamFlavour {
	// 	fmt.Println(p1.first, "Favourite icecream flavours are", v)
	// }

	// for _, v := range p2.favIcecreamFlavour {
	// 	fmt.Println(p2.first, "Favourite icecream flavours are", v)
	// }
	// fmt.Println("----------------------------------------------")
	m := map[string]person {
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIcecreamFlavour {
			fmt.Println(v.first, v.last, v2)
		}
	}
}


/*
Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the
slice.
https://go.dev/play/p/-9q96CysQfG
curriculum item # 127-hands-on-exercise-54-map-struct
*/