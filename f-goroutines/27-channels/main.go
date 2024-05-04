package main

import "fmt"

func main() {
	messages := make(chan string)

	for range 10 {
		go func() { messages <- "ping" }()
	}

	// blocks until one message is recieved
	msg := <-messages
	fmt.Println(msg)
}
