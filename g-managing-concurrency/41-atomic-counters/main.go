package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint32
	var wg sync.WaitGroup

	for range 5000 {
		wg.Add(1)

		go func() {
			for c := 0; c < 10000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
