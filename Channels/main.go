// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // func worker(ch chan string) {
// // 	fmt.Println("Worker waiting for data...")
// // 	msg := <-ch // Blocks until data is received
// // 	fmt.Println("Worker received:", msg)
// // }

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu.Lock() // Lock the critical section
// 	counter++ // Modify shared resource
// 	fmt.Printf("Worker %d incremented counter to %d\n", id, counter)
// 	mu.Unlock() // Unlock after modification
// }

// var counter int
// var mu sync.Mutex // Mutex for synchronization

// func main() {
// 	// ch:=make(chan int)
// 	// ch<-3
// 	// x:=<-ch
// 	// fmt.Println(x)

// 	// ch := make(chan string) // Unbuffered channel

// 	// go worker(ch)

// 	// time.Sleep(time.Second) // Simulating delay
// 	// ch <- "Hello"           // This will block until worker receives the message
// 	// time.Sleep(time.Second) // Simulating delay

// 	// fmt.Println("Main function finished")

// 	// ch := make(chan int, 3)
// 	// ch <- 1
// 	// ch <- 1
// 	// ch <- 1
// 	// // ch<-1
// 	// fmt.Println(<-ch)
// 	// ch <- 1
// 	// ch <- 1

// 	//mutex lock
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Final counter value:", counter)

// }

package main

import "fmt"

func test(ch chan int) {
	ch <- 2
	close(ch)
}
func test1(ch chan int) {
	ch <- 4
}

func main() {
	ch := make(chan int)
	go test(ch)
	go test1(ch)

	for ele := range ch {
		fmt.Println(ele)
	}
}
