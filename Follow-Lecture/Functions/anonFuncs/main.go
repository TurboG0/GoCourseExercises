package main

import (
	"fmt"
)

//named function
func foo() {
	fmt.Println("foo ran")
}

func main() {
	foo()

	//anonnymous function
	func(){
		fmt.Println("Hi this is anon func")
	}()
	
	func(s string){
		fmt.Println("Hi this is anon func which argument ", s)
	}("Gela")

	//func expression
	x := func() {
		fmt.Println("func expression ran")
	}
	//func expression
	y := func(s string){
		fmt.Println("Func expression with name", s)
		}
	x()
	y("Saba")
}


// a named function with an identifier
// func (r receover) identifier(p parameter(s)) {r returns(s)} {code}

// an anonymouse function
// func(p parametr(s)) (r returns(s)) {code}