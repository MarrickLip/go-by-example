package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name] += 1
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(id int, name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
			fmt.Println(id, c.counters)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement(1, "a", 50)
	go doIncrement(2, "a", 50)
	go doIncrement(3, "b", 50)

	wg.Wait()
	fmt.Println(c.counters)
}
