package main

import (
	"testing"
	"time"
)

func TestDataRace(t *testing.T) {
	input := make(chan string)
	go publisher(input)
	go publisher(input)
	go subscribeProcess(input)
	go subscribeProcess(input)
	time.Sleep(1 * time.Millisecond)
}
