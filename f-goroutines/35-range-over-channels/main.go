package main

func main() {
	queue := make(chan int, 5)

	for i := range 5 {
		queue <- i
	}
	close(queue)

	for elem := range queue {
		println(elem)
	}

}
