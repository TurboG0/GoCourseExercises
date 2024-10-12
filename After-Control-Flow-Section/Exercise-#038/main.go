package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	
	for k, v := range m {
		fmt.Printf("Key: %v \t Value: %v\n", k, v)
	}

	age := m["James"]
	fmt.Printf("Age of James is %v\n", age)

	if v, ok := m["James"]; ok {
		fmt.Printf("There is James look up entry, James age is %v\n", v)
	}

	age1 := m["Q"]
	fmt.Printf("age of Q is %v\n", age1)
	
	if v, ok := m["Q"]; !ok {
		fmt.Printf("There is no Q, and here is the zero value of an int %v\n", v)
	}

	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration: %v, x = %v\n", i, x)
		}
	}

}
