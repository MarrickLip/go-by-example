package main

// send-only channel: chan<- string
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings is read-only; pongs is send-only
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 2)
	pongs := make(chan string, 2)

	ping(pings, "hello")
	ping(pings, "world")

	pong(pings, pongs)
	pong(pings, pongs)

	println(<-pongs)
	println(<-pongs)
}
