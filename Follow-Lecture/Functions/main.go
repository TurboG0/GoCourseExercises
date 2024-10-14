package main

import (
	"fmt"
)
//multiple return
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years"), age
}
//variadic function that accept as many values as you want
// this int is turned into []int. T is turned into []T.
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

//receiver stuff

type person struct {
	first	string
}

type secretAgent struct {
	person
	ltk bool
}

//method for person
func (p person) speak() {
	fmt.Println("I am,", p.first)
}

//method for secretAgent
func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)	
}

//interface
//interface allows us to have polymorphism
type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}



func main() {
	s, n := dogYears("Saba", 22)
	fmt.Println(s, n)

	sliceForSum := []int {1, 2, 3, 4, 5}
	//sum accepts both arguments like this
	fmt.Println(sum(1, 2, 3, 4, 5))
	//... in this case opens up slice and passes individual values of slice
	// into the function
	fmt.Println(sum(sliceForSum...))

	p1 := person {
		first: "Saba", 
	}

	p2 := person {
		first: "Turbo", 
	}

	//p1.speak()
	//p2.speak()

	sa := secretAgent{
		person: person{
			first: "Saba-069",
		},
		ltk: true,
	}
	//sa.speak()
	saySomething(sa) 
	saySomething(p1) 
	saySomething(p2) 
}

//func (r receiver) identifier(p parameter(s)) {return(s)} {code}