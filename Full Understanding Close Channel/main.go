package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			ch <- i
		}
		fmt.Println("TOP")
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	time.Sleep(2 * time.Second)
	// for value := range ch {
	// 	fmt.Println(value)
	// }

	fmt.Println("fja")
}
