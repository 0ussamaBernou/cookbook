package main

import (
	"fmt"
	"time"
)

type Server struct {
	msgCh  chan Message
	quitCH chan struct{} // zero cost channel literally 0bytes
}

type Message struct {
	from    string
	payload string
}

func (s *Server) startAndListen() {
free:
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("received message: %s from %s\n", msg.payload, msg.from)

		case <-s.quitCH:
			fmt.Println("Server shutting down gracefully")
			break free
		default:
			// fmt.Println("waiting")
		}
	}
	fmt.Println("Server shut down")
}

func sendMessage(msgCh chan Message, payload string) {
	msg := Message{
		from:    "me",
		payload: payload,
	}
	msgCh <- msg
}
func main() {

	s := &Server{
		msgCh:  make(chan Message),
		quitCH: make(chan struct{}),
	}

	go s.startAndListen()

	go func() {
		time.Sleep(time.Second)
		sendMessage(s.msgCh, "Hello msgCh")
	}()
	go func() {
		time.Sleep(2 * time.Second)
		close(s.quitCH)
	}()
	select {}
}
