package concurrency

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello from goroutine")
		time.Sleep(500 * time.Millisecond)
	}
}

func Goroutines() {
	go sayHello() // starts a new goroutine
	for i := 0; i < 3; i++ {
		fmt.Println("Hello from main")
		time.Sleep(500 * time.Millisecond)
	}
}
