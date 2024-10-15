package main

import "fmt"

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}


func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)
	x := doMath(42, 16, add)
	fmt.Println(x)
	y := doMath(42, 16, sub)
	fmt.Println(y)
}

/*
Hands-on exercise #64 - tests in go #2 - unit tests
A unit test is a type of software testing where individual components or units of a
software are tested. The purpose is to validate that each unit of the software performs as
designed. A unit is the smallest testable part of any software. It usually has one or a few
inputs and usually a single output.
In Go programming language, a unit test usually tests a single function, method, or struct. The
goal is to confirm that the behavior is correct.
Unit tests in Go are typically written using the built-in testing package, `testing`. This
package does not require any third-party libraries, but it's somewhat limited compared to
some other languages' testing tools. Despite its simplicity, it has enough features to construct
effective unit tests.
When you ask if this differs from a "test" in Go, it's worth clarifying that a "unit test" is a subset
of "test". Tests in software can take many forms such as unit tests, integration tests,
functional tests, system tests, etc.
An "Integration Test", for example, in contrast to a "Unit Test", would test the interaction
between multiple components of the system, to ensure they work together correctly.
So, to summarize, in Go, a unit test is just one kind of test you can conduct, focused on
verifying the correct behavior of a small, isolated piece of functionality. Other types of tests
have different focuses and may require different tools or techniques.
create a unit test for these three functions
package main
import "fmt"
func main() {
fmt.Printf("%T\n", add)
fmt.Printf("%T\n", subtract)
fmt.Printf("%T\n", doMath)
x := doMath(42, 16, add)
fmt.Println(x)
y := doMath(42, 16, subtract)
fmt.Println(y)
}
Todd McLeod - Learn To Code Go - Page 102
func doMath(a int, b int, f func(int, int) int) int {
return f(a, b)
}
func add(a int, b int) int {
return a + b
}
func subtract(a int, b int) int {
return a - b
}
https://go.dev/play/p/ysOWyO9N3Kj
curriculum item # 157-hands-on-64-tests-02-unit-tests
*/
