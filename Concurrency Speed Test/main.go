package main

import (
	"fmt"
	"sync"
)

func Read(wg *sync.WaitGroup, winder int) {
	var counter int = 10

	fmt.Printf("======================== WINER %d ==================", winder)

	for index := 0; index < counter; index++ {
		fmt.Println(index)
	}

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	var loop int = 10

	for index := 0; index < loop; index++ {
		wg.Add(2)
		go Read(&wg, 1)
		go Read(&wg, 2)
	}

	wg.Wait()
}
