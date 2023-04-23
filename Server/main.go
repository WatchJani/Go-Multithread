package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	payload string
}

type Server struct {
	MessageCh chan Message
}

func New() *Server {
	return &Server{
		MessageCh: make(chan Message),
	}
}

func (s *Server) Listener() {
	for {
		select {
		case b := <-s.MessageCh:
			fmt.Println(b.From, b.payload)
		default:
		}
	}
}

func SendMessageToServer(ch chan Message, payload string) {
	msg := Message{
		From:    "Janko",
		payload: payload,
	}

	ch <- msg
}

func main() {
	s := New()
	go s.Listener()

	go func() {
		time.Sleep(2 * time.Second)
		SendMessageToServer(s.MessageCh, "SendMessageToServer")
	}()

	select {}
}
