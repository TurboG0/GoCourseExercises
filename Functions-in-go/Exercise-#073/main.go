package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

/*
Hands-on exercise #73 - wrapper
Here's a simple example to illustrate the concept of a wrapper function in Go:
package main
import (
"fmt"
"time"
)
// Wrapper function for adding timing information
func TimedFunction(fn func()) {
start := time.Now()
fn()
elapsed := time.Since(start)
fmt.Println("Elapsed time:", elapsed)
}
// Function to be wrapped
func MyFunction() {
time.Sleep(2 * time.Second) // Simulate some work
fmt.Println("MyFunction completed")
}
func main() {
// Call the wrapped function
TimedFunction(MyFunction)
}
In the example above, the `TimedFunction` acts as a wrapper function that measures the
elapsed time taken by `MyFunction` to execute. It captures the start time, calls `MyFunction`,
calculates the elapsed time, and then prints it. By using the wrapper, you can easily add
timing functionality to multiple functions without modifying their implementation.
Todd McLeod - Learn To Code Go - Page 109
Please note that the specific implementation of wrapper functions can vary depending on the
use case and requirements. Wrappers can be written as higher-order functions that accept
function arguments, or they can be implemented as methods of custom types. The choice of
implementation depends on the specific needs of your application.
A wrapper function, also known as a wrapper, is a function that encapsulates or wraps
another function or piece of functionality. It is typically used to provide additional behavior or
modify the behavior of the wrapped function without directly modifying its code. The wrapper
function takes the same arguments and returns the same type as the wrapped function, but it
may perform some pre- or post-processing, error handling, logging, or other tasks around the
execution of the wrapped function.
Here's an example of a simple wrapper function in Go:
```go
func Logger(f func()) {
fmt.Println("Starting execution...")
f()
fmt.Println("Execution completed.")
}
func main() {
wrappedFunc := func() {
fmt.Println("Hello, World!")
}
Logger(wrappedFunc)
}
```
In the above example, the `Logger` function is a wrapper that adds logging statements before
and after the execution of the wrapped function `wrappedFunc`.
On the other hand, a callback function is a function that is passed as an argument to another
function and is invoked by that function at a specific point or event. The purpose of a callback
function is to allow the receiving function to execute the callback code when appropriate or
necessary. Callback functions are commonly used in event-driven programming or
asynchronous operations.
Here's a simple example of using a callback function in Go:
Todd McLeod - Learn To Code Go - Page 110
```go
func ProcessData(data []int, callback func(int)) {
for _, item := range data {
callback(item)
}
}
func main() {
data := []int{1, 2, 3, 4, 5}
callbackFunc := func(num int) {
fmt.Println("Processing number:", num)
}
ProcessData(data, callbackFunc)
}
```
In the above example, the `ProcessData` function takes a slice of integers and a callback
function. It iterates over the data and invokes the callback function for each item. In the `main`
function, we define the callback function as an anonymous function that simply prints the
number being processed.
To summarize, a wrapper function wraps or modifies another function's behavior, while a
callback function is a function passed as an argument to be executed at a specific point or
event.
https://go.dev/play/p/55s4mfHLMdz
*/