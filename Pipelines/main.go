//SIMPLE PIPELINES
// package main

// import "fmt"

// // for producing the numbers
// // nums ...int - slice of int
// // will put the value in the channel
// func generator(nums ...int) <-chan int {
// 	out := make(chan int)

// 	go func() {
// 		for _, val := range nums {
// 			out <- val
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func square(nums <-chan int) <-chan int {
// 	result := make(chan int)

// 	go func() {
// 		for val := range nums {
// 			result <- val * 2
// 		}
// 		close(result)
// 	}()

// 	return result
// }

// func main() {
// 	nums := generator(1, 2, 3, 4, 5)
// 	result := square(nums)

// 	for val := range result {
// 		fmt.Println(val)
// 	}

// }

// FAN OUT - single channel se multiple go routines take data
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Printf("%d job is processed by worker-%d \n", job, id)
// 		result <- job * 2
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	jobs := make(chan int)
// 	results := make(chan int)
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}

// 	go func() {
// 		for i := 0; i < 8; i++ {
// 			jobs <- i
// 		}
// 		close(jobs)
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()
// 	for res := range results {
// 		fmt.Println(res)
// 	}

// }

//FAN IN - result from multiple channel into one channel

package main

import (
	"fmt"
	"sync"
)

func fanIN(ch1 <-chan int, ch2 <-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(2)
	result := make(chan int)

	go func() {
		for val1 := range ch1 {
			result <- val1
		}
		wg.Done()
	}()

	go func() {
		for val2 := range ch2 {
			result <- val2
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	return result

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	result := fanIN(ch1, ch2)

	for val := range result {
		fmt.Println(val)
	}

}
