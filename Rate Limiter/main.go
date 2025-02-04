package main

import (
	"fmt"
	"time"
)

// RateLimiter struct
type RateLimiter struct {
	tokens chan struct{}
}

// NewRateLimiter creates a new rate limiter with a fixed rate
func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, rate),
	}

	// Fill the token bucket at a fixed rate
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate)) // Add a token every interval
		defer ticker.Stop()

		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}: // Add a token if the bucket is not full
				fmt.Printf("Token added at: %s\n", time.Now().Format(time.RFC3339))
			default:
			}
		}
	}()

	return rl
}

// Allow checks if a request can proceed
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true // Allowed
	default:
		return false // Rate limited
	}
}

func main() {
	rate := 2                       // 2 requests per second
	limiter := NewRateLimiter(rate) // Create rate limiter

	// Simulating 10 requests
	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: Processed \n", i)
		} else {
			fmt.Printf("Request %d: Rate limited \n", i)
		}
		time.Sleep(300 * time.Millisecond) // Simulating incoming requests
	}
}
