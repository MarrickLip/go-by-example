package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Println(">>", id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println("<<", id)
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
