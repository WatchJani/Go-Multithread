package main

import "fmt"

func Gen(list []string) <-chan string {
	ch := make(chan string)

	go func() {
		for index := range list {
			ch <- list[index]
		}
		close(ch)
	}()

	return ch
}

func Prefix(ch <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for word := range ch {
			out <- "GO_" + word + ".txt"
		}
		close(out)
	}()

	return out
}

func Reader(ch <-chan string) {
	for newWord := range ch {
		fmt.Println(newWord)
	}
}

func main() {
	textEdit := []string{"Test", "Golang", "Pipeline", "Channel"}

	newCh := Gen(textEdit)

	Reader(Prefix(newCh))
}
