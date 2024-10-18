package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	//throtteling
	const goroutines = 10
	wg.Add(goroutines)
	
	//throtteling
	for i := 0; i < goroutines; i++ {
		for v := range c1 {
			go func(v2 int) {
				c2 <- timeConsumptingWork(v2)
				wg.Done()
			}(v)
		}

	}
	wg.Wait()
	close(c2)
}

func timeConsumptingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
