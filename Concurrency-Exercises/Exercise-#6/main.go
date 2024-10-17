package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	time.Sleep(time.Second * 10)
}

/*
Hands-on exercise #6
Create a program that prints out your OS and ARCH. Use the following commands to run it
● go run
● go build
● go install
code: https://github.com/GoesToEleven/go-programming
*/
