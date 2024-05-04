package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := range 5 {
		requests <- i
	}
	close(requests)

	// a basic limter waits on the ticker
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// a bursty limiter uses a buffered channel

	// burstyLimiter := make(chan int, 3)

}
