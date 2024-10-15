package main

import "fmt"

func main() {
	p := Person {
		first:	"Saba",
		age:	22,
	}
	p.Speak()
}

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Println("My name is",p.first, "and my age is",p.age)
}

/*
Hands-on exercise #61 - method
● Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first
■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
https://go.dev/play/p/KqOV32q1aC0

*/