package main

import (
	"fmt"
)

func main() {
	am := map[string]int {
		"Saba": 22,	
		"Lui": 3,
		"Noa": 4,
		"Lea": 2,
		"Jefi": 6,
	}

	fmt.Println("The age of Saba is ", am["Saba"])

	//print whole map
	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	fmt.Println("------------------------------")
	//creating map using make
	an := make(map[string]int)
	an["Gela"] = 37
	an["Bichiko"] = 40
	an["Dildo Beggins"] = 69
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	//length of map
	fmt.Println("Length", len(an))
	
	
	//use for range to loop over map
	fmt.Println("------------------------------")

	for k, v := range an {
		fmt.Println(k, v)
	}
	//if you only want value
	for _, v := range an {
		fmt.Println(v)
	}
	
	//if you only want key
	for k := range an {
		fmt.Println(k)
	}

	//map delete element
	fmt.Println("------------------------------")

	delete(an, "Gela") 
	// won' panic even though element is already deleted if it existed 
	delete(an, "Gela")

	fmt.Println(an)

	fmt.Println("------------------------------")
	//accesing key that is already deleted
	delete(an, "Saba") // won't panic
	fmt.Println(an["Saba"]) // won't panic
	
	//comma ok idiom
	fmt.Println("------------------------------")
	
}