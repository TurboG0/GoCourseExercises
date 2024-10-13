package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-----------------")
	
	// [Inclusive:Exclusive]
	fmt.Printf("xi - %#v\n", xi[0:4])
	fmt.Println("-----------------")
	
	// [:Exclusive]
	fmt.Printf("xi - %#v\n", xi[:7])
	fmt.Println("-----------------")
	
	//["inclusive:"]
	fmt.Printf("xi - %#v\n", xi[4:])
	fmt.Println("-----------------")
	
	//[:]
	fmt.Printf("xi - %#v\n", xi[:])
	fmt.Println("-----------------")

	//deleting from slice
	//xi = append(xi[:4], xi[5:]...)
	fmt.Println(xi)
	fmt.Println("-----------------")
	xi = append(xi[:0], xi[5:]...)
	fmt.Println(xi)

	//making slice with make keywoard
	//first argument is type second is how many elements we store right away in slice
	//third argument is initial size of the slice
	//if we skip 2nd argument it will default to 0
	fmt.Println("-----------------")
	s1 := make([]int, 0, 10)
	
	fmt.Printf("Slice: %v\n",s1)
	fmt.Printf("Lenght of slice: %v \n", len(s1))
	fmt.Printf("Capacity of slice: %v\n", cap(s1))
	s1 = append(s1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("Slice after append: %v\n",s1)
	
	fmt.Println("-----------------")	
	
	s1 = append(s1, 10, 11, 12, 13, 14, 15, 16)
	fmt.Printf("Slice: %v\n",s1)
	fmt.Printf("Lenght of slice: %v \n", len(s1))
	fmt.Printf("Capacity of slice: %v\n", cap(s1))
	fmt.Printf("Slice after 2nd append: %v\n",s1)
	fmt.Println("-----------------")	
	
	//proving that Slice is pointing to underlaying array and is just a point
	a := []int{1, 2, 3, 4}
	b := a
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-----------------")
	//This will also change b
	a[0] = 7
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-----------------")
	//Copying 1 slice to other and not pointing
	x := make([]int, 6)
	copy(x, a)
	a[1] = 19
	fmt.Println("a", a)
	fmt.Println("x", x)
	fmt.Println("-----------------")
	
	n := []float64{3, 1, 4, 2}
	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n)
	
	n2 := []float64{3, 1, 4}
	fmt.Println(medianTwo(n2))
	fmt.Println("after medianTwo", n2)
}

// uses the same underlying array
// everything is pass by value in go
// the vaulue is being passed into the func
// and then assigned to x

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x) % 2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2
	if len(n) % 2 == 1 {
		return n[i/2]
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
	}
	return (n[i - 1] + n[i]) / 2
}