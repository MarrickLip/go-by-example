package main

import "fmt"

func main() {
	// messages is buffered, allowing it to queue two messages
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
