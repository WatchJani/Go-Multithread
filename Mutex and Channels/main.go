package main

import (
	"fmt"
	"sync"
)

// Kada se locka st struktura, svi njeni polja su lockani zajedno.
// U ovom slučaju, nakon što se locka struktura, kôd pokušava pročitati iz ch kanala.
// No, može se dogoditi da se drugi go-routine, koja piše u kanal, pokuša izvršiti upis prije nego što se lock oslobodi.
// To će dovesti do blokade (deadlock) koda, jer su i upis i čitanje kanala lockani zajedno,
// te se ne može izvršiti upis u kanal dok je njegovo čitanje blokirano.

// type Example struct {
// 	ch      chan int64
// 	counter int64
// 	looping int
// 	sync.WaitGroup
// 	sync.RWMutex
// }

// func New() *Example {
// 	return &Example{
// 		ch:      make(chan int64), //change nil value
// 		looping: 10,
// 	}
// }

type Example struct {
	counter int64
	looping int
	// sync.WaitGroup
	sync.RWMutex
}

func New() *Example {
	return &Example{
		looping: 10,
		counter: 1,
	}
}

func Set(st *Example, ch chan int64) {
	st.Lock()
	defer st.Unlock()

	ch <- st.counter
	st.counter++
}

func Reader(st *Example, ch chan int64) {
	for i := 0; i < st.looping; i++ { //here is going infinite loop if we use WaitGroup
		fmt.Println(<-ch)
		// st.Done()
	}
}

func main() {
	st := New()
	ch := make(chan int64)

	// st.Add(st.looping)
	for i := 0; i < st.looping; i++ {
		go Set(st, ch)
	}

	// go Reader(st, ch)
	// st.Wait()

	Reader(st, ch)
}
