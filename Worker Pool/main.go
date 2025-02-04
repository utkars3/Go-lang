package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that processes jobs from the jobs channel
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Processing result
	}
}

func main() {
	numWorkers := 2 // Number of workers
	numJobs := 5    // Number of jobs

	jobs := make(chan int, numJobs)    // Channel for jobs
	results := make(chan int, numJobs) // Channel for results

	var wg sync.WaitGroup

	// Create worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close channel after sending all jobs

	// Wait for workers to complete
	wg.Wait()
	close(results) // Close results channel

	// Collect and print results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
