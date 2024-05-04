package main

import (
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
		}
	}
}
