package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("file.txt")

	if err != nil {
		fmt.Println("error")
		return
	}

	defer f.Close()

	f.Write([]byte("ზდაროვა ბრატცი კეკწ 😂😂"))

	bs, err := io.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

// always check your god damn errors.
