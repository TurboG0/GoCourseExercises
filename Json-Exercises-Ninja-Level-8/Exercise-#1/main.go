package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string
	Last  string
	Age   int
}

func main() {
	u1 := user{
		First: "Saba",
		Last:  "Dzodzuashvili",
		Age:   22,
	}

	u2 := user{
		First: "Luke",
		Last:  "Skywanker",
		Age:   69,
	}

	users := []user{
		u1,
		u2,
	}

	bs, err := json.Marshal(users)

	if err != nil {
		log.Fatalf("Error - %v\n", err)
	}

	s := string(bs)

	fmt.Println(s)
}

//[{"First":"Saba","Last":"Dzodzuashvili","Age":22},
//{"First":"Luke","Last":"Skywanker","Age":69}]
/*
Hands-on exercise #1
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this
hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
package.
solution: https://play.golang.org/p/8BK6PXj3aG
video: 125
*/
