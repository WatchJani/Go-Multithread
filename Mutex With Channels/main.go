package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var y int

func Add(ch chan struct{}, wg *sync.WaitGroup) {
	ch <- struct{}{}
	y += 50
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan struct{}, 1)

	for index := 0; index < rand.Intn(50); index++ {
		wg.Add(1)
		go Add(ch, &wg)
	}

	wg.Wait()

	fmt.Println(y)
}
