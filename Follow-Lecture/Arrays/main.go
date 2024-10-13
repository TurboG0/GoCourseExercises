package main 

import (
	"fmt"
)

func main() {
	a := [3]int {42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Saba"}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4}
	fmt.Println(c)

	var d [2] int
	d[0] = 5
	d[1] = 6
	fmt.Println(d)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)

}