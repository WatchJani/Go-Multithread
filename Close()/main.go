package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(time.Minute)
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	fmt.Println("top")
}
