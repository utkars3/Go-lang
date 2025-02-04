package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func prnt2(id int) {
	fmt.Println("doing work id vh", id)
}

func work(id int, wg *sync.WaitGroup) {
	once.Do(func() { prnt2(id) })
	// prnt(id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	fmt.Println("work completed")

}
