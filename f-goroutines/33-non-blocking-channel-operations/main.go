package main

import "fmt"

func main() {
	messages := make(chan string)

	// non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("no messages recieved")
	}

	msg := "hello world"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message ssend")
	}

}
