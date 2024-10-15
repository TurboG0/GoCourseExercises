package main

import "fmt"

type person struct {
	first string
}

func changeName(p person, name string) person {
	p.first = name
	return p
}

func changeNameP(p *person, s string) {
	p.first = s
}

func main() {
	p := person{"Saba"}
	p = changeName(p, "Gela")
	fmt.Println(p.first)
	changeNameP(&p, "Bayayi")
	fmt.Println(p.first)
}

/*
Hands-on exercise #77 - value & pointer semantics
Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS
○ this function will return nothing
● don't use methods
HINT:
In Go, a value of type `T` can use a method with a receiver of type `*T` but with some
limitations.
If the method has a pointer receiver, `*T`, it means it needs to modify the receiver or the
receiver is a large struct that would be expensive to copy.
In this case, if you try to call this method using a value of type `T` (not `*T`), Go will
automatically take the address of that value and use that pointer as the receiver. However,
it is important to note that this is possible only if the value of type `T` is addressable, i.e., it is
Todd McLeod - Learn To Code Go - Page 123
a variable, not an unaddressable value like a literal or a temporary result of a function or
expression.
Here is a small example:
```go
type T struct {
name string
}
func (t *T) changeName(newName string) {
t.name = newName
}
func main() {
var myT T
myT.changeName("new name") // Works fine, Go automatically takes the address of myT
fmt.Println(myT.name) // Prints "new name"
T{"another name"}.changeName("new name") // Compiler error, T{"old name"} is not
addressable
}
```
In the second example, `T{"another name"}.changeName("new name")`, the compiler would
throw an error since `T{"another name"}` is not an addressable value. Therefore, a method
with a pointer receiver cannot be called on it.
So while a value of type `T` can use a method with a receiver of type `*T`, this only applies to
addressable values of type `T`. If the value isn't addressable, you'll get a compile-time error.
FYI, certain things are addressable and certain things are not. Generally, you can take the
address of variables, elements of an array or slice, fields of an addressable struct, and you
can't take the address of constants, literals (like `T{"another name"}`), or the return values of
functions or expressions (unless they return a pointer type or an interface type).
The reason that the struct literal `T{"another name"}` is not addressable is because it's a
temporary value in memory. It's not stored in a variable or a field, it's just a transient piece of
data that's created on the fly and discarded immediately after. It's not allocated a stable
location in memory that you can reference, hence it doesn't have an address you can take.
If you want to call a method that has a pointer receiver on a struct, you can create a pointer to
the literal directly:
```go
(&T{"another name"}).changeName("new name")
```
Todd McLeod - Learn To Code Go - Page 124
In this case, `&T{"another name"}` creates a pointer to the struct literal and you can call
methods with a pointer receiver on it.
But remember that any changes made inside `changeName` method will not persist outside
the scope of the method call because the original `T` value is a temporary value. This is
different from the case when you have an addressable variable where the changes made by
the method will persist as they're modifying the actual memory location where the variable is
stored.
GUIDELINES - VALUE SEMANTICS
"When we're working with builtin types – when the data we're working with is either numeric,
string, or bool – what we're going to want to use is our value semantics to move that data
around the program." - Bill Kennedy
There are three types in go: builtin, reference, and structs (user defined type). The zero value
of a reference type is nil
https://go.dev/play/p/TDnuZTYbHdp
*/