// 5). Write a program that launches a goroutine to listen to a channel for incoming strings.
// Whenever a string is received, the goroutine should print out the string with a timestamp.
// The program should then send 10 strings to the channel and terminate.

package helpers

import (
	"fmt"
	"sync"
)

type Message struct {
	Channel chan string
}

func NewMessage() *Message {
	return &Message{
		Channel: make(chan string),
	}
}

var mc = NewMessage()

func WriteMessage(ws *sync.WaitGroup) {
	go func() {
		defer ws.Done()
		for i := 0; i < 10; i++ {
			mc.Channel <- fmt.Sprintln("Hello world", i)
		}
		close(mc.Channel)
	}()
}

func ReadMessage(ws *sync.WaitGroup) {
	go func() {
		defer ws.Done()
		for {
			if chanVal, ok := <-mc.Channel; ok {
				fmt.Println("Value received from channel is", chanVal)
			} else {
				fmt.Println("Channel Terminated")
				break
			}
		}

	}()
}
