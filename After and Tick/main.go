package main

import (
	"fmt"
	"time"
)

func Create() <-chan struct{} {
	ch := make(chan struct{})

	time.Sleep(time.Second)

	// go func() {
	// }()
	// close(ch)

	return ch
}

func main() {
	ch := Create()

	select {
	case <-ch:
		fmt.Println("top")
		// default:
		// 	fmt.Println("waith")
		// 	time.Sleep(100 * time.Millisecond)
	}
}
