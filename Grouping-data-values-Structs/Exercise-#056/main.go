package main

import (
	"fmt"
)


func main() {
	anon := struct {
		first		string
		friends		map[string]int
		favDrinks	[]string
	}{
		first:		"Saba",
		friends:	map[string]int {
			"Shalia":	21,
			"Givi":		22,
			"Beggins":	489,
		},
		favDrinks:  []string {"kakao, kola, wyali, yava, chai"},
	}
	
	for k, v := range anon.friends{
		fmt.Println(anon.first, "friends - ", k, v)
	}
	
	for _, v := range anon.favDrinks{
		fmt.Println(anon.first, "favourite drinks - ", v)
	}
}

/*
Create and use an anonymous struct with these fields:
● first string
● friends map[string]int
● favDrinks []string
Print things.
https://go.dev/play/p/l72ejdKEe3r
curriculum item # 129-hands-on-exercise-56-anon-struct
*/