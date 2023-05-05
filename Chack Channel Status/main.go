package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	time.Sleep(time.Second)
	close(ch)

	value, ok := <-ch

	fmt.Println(value)

	fmt.Println(ok)
}
