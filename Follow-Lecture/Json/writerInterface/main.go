package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, bro")
	fmt.Fprintln(os.Stdout, "Hello, bro")
	io.WriteString(os.Stdout, "Hello, bro")
	
}
