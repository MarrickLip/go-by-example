package main

import (
	"fmt"
	"time"
)

func f(prefix string) {
	for i := range 100 {
		println(prefix, i)
	}
}

func main() {
	// run synchronously
	f("direct")

	// run asynchronously
	go f("async")

	go func() {
		fmt.Println("anonymous go routine")
	}()

	time.Sleep(time.Second)
	fmt.Println("done")
}
