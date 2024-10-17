package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p *person) speak() {
	fmt.Printf("Hello, my full name is %v %v and I am %v years old. I can speak\n", p.First, p.Last, p.Age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// can
	p := person {
		First: "Saba",
		Last: "Dzodzuashvili",
		Age: 22,
	}
	// can't do this because speak() has *T Receiver
	// saySomething(p)
	saySomething(&p)
}

/*
Hands-on exercise #2
This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
○ *person
● create a type human interface
○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
○ have it take in a human as a parameter
○ have it call the speak method
● show the following in your code
○ you CAN pass a value of type *person into saySomething
○ you CANNOT pass a value of type person into saySomething
● here is a hint if you need some help
○ https://play.golang.org/p/FAwcQbNtMG
Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T
code: https://github.com/GoesToEleven/go-programming
*/
