package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working")

	for range 3 + rand.Intn(6) {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
