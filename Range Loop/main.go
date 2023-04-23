package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Push(ch chan int) {
	timer := rand.Intn(3) + 1
	time.Sleep(time.Duration(timer) * time.Second)

	ch <- timer
}

//izgleda da nece blokirati
func main() {
	ch := make(chan int)

	for index := 0; index < 10; index++ {
		go Push(ch)
	}

	close(ch)
	time.Sleep(4 * time.Second)

	for value := range ch {
		fmt.Println(value)
	}
}
