package main

import (
	"fmt"
	"time"
)

func CreateNewTimer(duration time.Duration) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		for {
			time.Sleep(duration * time.Millisecond)
			ch <- struct{}{}
		}
	}()

	return ch
}

func main() {
	timer := CreateNewTimer(5000)

	for {
		select {
		case <-timer:
			fmt.Println("Potrosio si novih", 5, "s u zivotu")
		}
	}
}
