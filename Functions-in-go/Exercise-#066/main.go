package main

// provides printing function
import "fmt"

func main() {
	fmt.Println(vtestaob("Usashvelo"))
}
// vtestaob prints out user's argument for function argument
func vtestaob(arg string) string {
	return fmt.Sprint("Funqciis argumentia: ", arg)
}


/*
Hands-on exercise #66 - documenting code with comments
● to document
○ a type, variable, constant, function, or package,
○ write a comment directly preceding its declaration, with no intervening
blank line.
■ begin with the name of the element
■ for packages
● first sentence appears in package list
● if a large amount of documentation, place in its own file
doc.go
● document our code from the previous exercise
In Go programming, the convention for documenting code uses comments immediately
preceding the declaration or definition of the item being documented. This style of
documentation is used by the `godoc` tool to automatically generate API documentation.
1. Package documentation: This should be a brief one-liner that begins with "Package
`<packagename>`" and provides a high-level overview of the package's purpose. It is typically
located in a separate `doc.go` file.

```go
// Package math provides mathematical constants and functions.
package math
```
2. Function/method documentation: Comments should be written directly above the
function, starting with the function's name.
```go
// Add calculates the sum of two integers.
func Add(x, y int) int {
return x + y
}
```
3. Type documentation: Similar to functions, comments should be written directly above the
type definition, starting with the type's name.
```go
// Person represents a person with a name and age.
type Person struct {
Name string
Age int
}
```
4. Constant and variable documentation: Follows the same pattern, starting with the
constant or variable name.
```go
// Pi is the ratio of the circumference of a circle to its diameter.
const Pi = 3.14159265358979323846
```
Other Go comment conventions include:
- Use complete sentences. The first sentence should be a summary starting with the name
being declared.
- Prefer to break long lines after punctuation or operators, keeping operands together.
- Do not use URLs in package-level comments. It's better to create a README file for that
purpose.
Remember, the goal of comments is to improve code readability and provide context to other
developers. It's not about explaining what the code does (the code should speak for itself in
that regard), but rather why it does it.
The go doc command shows documentation for exported names (i.e., identifiers that start with
Todd McLeod - Learn To Code Go - Page 104
a capital letter). Unexported identifiers (i.e., those that start with a lowercase letter) are private
to their package and are not included in the output of go doc by default.
However, if you want to see documentation for unexported identifiers in your package while
you're developing, you can use go doc -u. This will show the documentation for both exported
and unexported identifiers. Note that these options are useful for package developers but are
less useful for consumers of the package since they cannot access unexported identifiers.
Please check the current Go documentation or help for go doc (go help doc) for the most
up-to-date information, as the behavior and features of Go tools may change.
https://go.dev/play/p/jxFzGsHyHfQ
*/