package main

import "fmt"

type dog struct {
	first string
}

// value receiver
func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

// pointer receiver
func (d *dog) run() {
	d.first = "Maximus"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	// value semantics
	d1 := dog{"Jefi"}
	d1.walk()
	d1.run()

	// using interface
	//youngin(d1) // this doesn't work because method run has pointer receiver


	// pointer semantics. pointer to dog
	d2 := &dog{"Maxi"}
	d2.walk()
	d2.run()
	// This works because when you have a value which is a pointer ,
	// you could use pointer receivers or value receivers to implement
	// and interface
	// if you have a value which is just value, not a pointer, you can
	// only use value receivers and this is why youngin(d1) doesn't work
	// and this works
	youngRun(d2)
}
