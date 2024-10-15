package main 

func main() {

}

func Add(a int, b int) int {
	return a + b
}

/*
Hands-on exercise #63 - tests in go #1
● Read the following
● run the code
● read through the code
● try to understand the code
Go provides a built-in testing framework that simplifies the process of testing Go code. Here is
an overview of the file structure, naming conventions, and how testing works in Go:
File Structure and Naming Conventions:
1. Test files: Test files live in the same package as the code they test. They are named with
the following convention: `filename_test.go` where filename is the name of the file that
contains the code you want to test.
2. Test functions: Test functions must start with the word `Test` followed by a word that starts
with a capital letter. The signature of a test function is `func TestXxx(*testing.T)`, where Xxx
does not start with a lowercase letter.
Example:
Todd McLeod - Learn To Code Go - Page 100
Suppose you have a function in `main.go` file that adds two integers:
```go
// main.go
package main
func Add(a int, b int) int {
return a + b
}
```
To test this function, you would create another file in the same directory called `main_test.go`.
```go
// main_test.go
package main
import "testing"
func TestAdd(t *testing.T) {
total := Add(5, 5)
if total != 10 {
t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
}
}
```
In the test function, we used the `Add` function and provided it with two integers. We then
compared the result with the expected value. If the result is not what we expect, we report an
error using `t.Errorf`.
Running the Tests:
To run tests in Go, open your terminal and navigate to the directory containing your files, then
type:
```shell
go test
```
If the test fails, the `go test` command will report an error and provide the output from the
`t.Errorf` function.
Go's testing package also provides more advanced features like benchmarks, test helpers,
and skipping tests. To explore more advanced usage, you can check the official Go testing
documentation.
https://go.dev/play/p/GhR7rzvxU3_p
*/