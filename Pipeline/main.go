package main

import (
	"fmt"
	"time"
)

func Gen() (<-chan int, int) {
	out, number_of_iteration := make(chan int), 10

	go func() {
		for i := 0; i < number_of_iteration; i++ {
			time.Sleep(time.Second)
			out <- i
		}
	}()

	return out, number_of_iteration
}

func Reader(ch <-chan int, num int) {
	for i := 0; i < num; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	Reader(Gen())
}
