package main

import (
	"fmt"
	"time"
	// "time"
)

func functionOne(chan1 chan string) {
	// time.Sleep(1 * time.Second)
	chan1 <- "Welcome to func 1"
	// chan1 <- "Welcome to func 1"
}

func functionTwo(chan2 chan string) {
	time.Sleep(1 * time.Second)
	chan2 <- "Welcome to func 2"
}

func main() {
	// chan1 := make(chan string)
	// chan2 := make(chan string)

	// go functionOne(chan1)
	// go functionTwo(chan2)

	// select {
	// case val1 := <-chan1:
	// 	fmt.Println(val1)
	// case val2 := <-chan2:
	// 	fmt.Println(val2)
	// 	// default:
	// 	// 	fmt.Println("Default")

	// }




	//deadlock cases
	// ch := make(chan int)

	// go func() {
	// 	for i := 1; i <= 3; i++ {
	// 		ch <- i
	// 	}
	// 	// Missing: close(ch) to terminate the range loop
	// 	close(ch)
	// }()

	// for num := range ch {
	// 	fmt.Println(num) // Stuck forever after receiving 3 values
	// }

}
