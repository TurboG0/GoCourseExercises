package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 3))
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

func square(n int) int {
	return n * n
}
/*
Hands-on exercise #71 - callback
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
○ func square(n int) int
○ printSquare(f func(int)int, int) string
Callback functions are very common in many programming languages including Go.
Callbacks are essentially functions that are passed as arguments to other functions and are
intended to be called (or "executed") later in your program.
Let's look at a few simple examples.
1. **A Basic Example**:
This first example shows a simple callback function. We're creating a function, `process`, that
takes a callback function as an argument. It then uses this callback to process an integer.
```go
package main
import "fmt"
// Our callback function type.
type callbackFunc func(int) int
// A function that takes a callback.
func process(cb callbackFunc, num int) int {
return cb(num)
}
Todd McLeod - Learn To Code Go - Page 106
// Our callback function.
func square(n int) int {
return n * n
}
func main() {
result := process(square, 5)
fmt.Println(result) // Output: 25
}
```
2. **Callback with Anonymous Function**:
This example uses an anonymous function (or lambda function) as a callback.
```go
package main
import "fmt"
type callbackFunc func(int) int
func process(cb callbackFunc, num int) int {
return cb(num)
}
func main() {
// Define an anonymous function that doubles the input.
doubler := func(n int) int {
return n * 2
}
result := process(doubler, 5)
fmt.Println(result) // Output: 10
}
```
3. **Callback in a Iteration**:
The third example iterates over a slice of integers, applying a callback function to each one.
```go
package main
import "fmt"
type callbackFunc func(int) int
Todd McLeod - Learn To Code Go - Page 107
func iterate(nums []int, cb callbackFunc) {
for i, num := range nums {
nums[i] = cb(num)
}
}
func main() {
nums := []int{1, 2, 3, 4, 5}
// Define a function that triples the input.
tripler := func(n int) int {
return n * 3
}
iterate(nums, tripler)
fmt.Println(nums) // Output: [3 6 9 12 15]
}
```
These examples demonstrate how to use callback functions in various contexts in Go, and
they should provide a good starting point for understanding this concept.
In Go, `type` is used to define new data types or to create an alias for an existing type.
When we write `type callbackFunc func(int) int`, we're defining a new type, `callbackFunc`,
which represents a function that takes an `int` as an argument and also returns an `int`.
This is why, in the `process` function, the `cb` parameter is of type `callbackFunc`. It's saying
that `cb` is a function that takes an `int` as an argument and returns an `int`.
So, when we pass the `square` function as the `cb` argument to the `process` function, Go
checks to see if `square` matches the `callbackFunc` type. Since `square` is defined as `func
square(n int) int` -- a function that takes an `int` as an argument and returns an `int` -- it
matches the `callbackFunc` type, and can be passed as the `cb` parameter to `process`.
This is essentially how types work in Go, and how we can define function types that can be
used as parameters for other functions. It provides a lot of flexibility and power when
structuring our programs, as it allows us to write functions that can take other functions as
arguments.
https://go.dev/play/p/o7klLEt0rEE
*/