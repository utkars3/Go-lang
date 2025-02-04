package main

import (
	"fmt"
	"time"
)

var counter int

// var mu sync.Mutex

func increment() {
	for i := 0; i < 100000; i++ {
		// mu.Lock()
		counter++
		// mu.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ { // Increase concurrent goroutines
		go increment()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(counter)

}
