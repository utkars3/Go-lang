// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func sayHello() {
// 	fmt.Println("Hello from Goroutine!")
// 	fmt.Println("Main function finished-2")

// }

// // The go keyword creates a new goroutine.
// // The sayHello() function runs independently from the main() function.
// // The main function doesn’t wait for sayHello() to complete.
// // A goroutine keeps running until it makes a blocking call, such as I/O, time.Sleep(), or runtime.Gosched().

// func task(id int) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("Task %d - iteration %d\n", id, i)
// 		runtime.Gosched() // Yield to other goroutines
// 	}
// }

// func main() {
// 	// go sayHello()           // Runs concurrently
// 	// fmt.Println("Main function finished-3")

// 	// time.Sleep(time.Second) // Give goroutine time to execute
// 	// fmt.Println("Main function finished")

// 	go task(1)
// 	go task(2)
// 	time.Sleep(time.Second)
// }

//go routines in depth

package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //marking done after completing the function will reduce the counter by 1
	fmt.Println("Task started ", i)
	fmt.Println("Task ended ", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		go worker(i, &wg)
		wg.Add(1) //adding 1 so that wg know that how many task are in go routine
	}
	wg.Wait() //it will make sure the program will not move forward until the wg counter will become zero
	fmt.Println("All task completed")
}
