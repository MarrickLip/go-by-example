package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		results <- j * id
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("worker", id, "finished job", j)
	}
}

func main() {
	numJobs := 5
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs*numWorkers)

	for w := range numWorkers {
		go worker(w, jobs, results)
	}

	for j := range numJobs {
		jobs <- j
	}

	close(jobs)

	for range numJobs {
		<-results
	}

}
