package main 

import (
	"fmt"
)

func main() {
	mySlice := []string {"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", 
	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
	"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
	"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", 
	"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", 
	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", 
	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", 
	"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", 
	"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", 
	"Wolverine Tracks (GF)"}

	fmt.Println(mySlice)
	fmt.Printf("Type of mySlice is: %T\n", mySlice)

	for i, v := range mySlice {
		fmt.Printf("index is: %v\tvalue is: %v\n", i, v)
	}
	fmt.Println(len(mySlice))
}


/*
Use a slice literal to store these elements in a slice, then also print out the slice and the
length of the slice.
● if you can, try to get a "for range" loop working on the slice
○ you can reference the documentation for a "for range" loop here:
■ https://go.dev/doc/effective_go#for
"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet
Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter
Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar
Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate
Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter
Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry
Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade
Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine
Tracks (GF)"
https://go.dev/play/p/W8sWn7fsUXh
curriculum item # 096-hands-on-exercise-41


*/