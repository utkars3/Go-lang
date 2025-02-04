// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context, id int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Worker %d canceled\n", id)
// 			return
// 		default:
// 			// Simulate work
// 			fmt.Printf("Worker %d is working...\n", id)
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	// Start workers
// 	go worker(ctx, 1)
// 	go worker(ctx, 2)

// 	// Run for 2 seconds and then cancel the context
// 	time.Sleep(2 * time.Second)
// 	cancel()

// 	// Wait for a moment to let workers finish
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("All workers stopped")
// }





//with cancel
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context, id int) {
// 	select {
// 	case <-time.After(3 * time.Second):
// 		// Simulate work
// 		fmt.Printf("Worker %d finished work\n", id)
// 	case <-ctx.Done():
// 		// Timeout or cancellation occurred
// 		fmt.Printf("Worker %d canceled: %v\n", id, ctx.Err())
// 	}
// }

// func main() {
// 	// Create a context with a 2-second timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	// Start a worker
// 	go worker(ctx, 1)

// 	// Wait for the worker to finish
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Main function finished")
// }





// with value
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	// Retrieve a value from the context
	user := ctx.Value("user")
	fmt.Printf("Worker %d: User from context: %v\n", id, user)
}

func main() {
	// Create a background context and add a value
	ctx := context.WithValue(context.Background(), "user", "Utkarsh")

	// Start workers
	go worker(ctx, 1)
	go worker(ctx, 2)

	// Wait before exiting
	time.Sleep(time.Millisecond)
	fmt.Println("Main function finished")
}
