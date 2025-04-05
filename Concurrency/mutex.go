package concurrency

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}

func Mutex() {
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go increment()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
