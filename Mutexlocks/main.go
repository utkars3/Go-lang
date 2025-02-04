package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	fmt.Printf("worker %d incremented counter to %d\n", id, counter)
	mu.Unlock()
}

func main() {
	fmt.Println("Working with mutex locks")
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value ", counter)

}
