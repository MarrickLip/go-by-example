package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		c1 <- "result"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

}
