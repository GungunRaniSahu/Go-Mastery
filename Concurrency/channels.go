package concurrency

import (
	"fmt"
)

func sendMessage(ch chan string) {
	ch <- "Hello from Goroutine!"
}

func Channels() {
	messageChannel := make(chan string)

	go sendMessage(messageChannel) 

	msg := <-messageChannel 
	fmt.Println(msg)
}
