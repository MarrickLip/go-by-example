package main

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for j, more := <-jobs; more; j, more = <-jobs {
			println("received job", j)
		}
		done <- true
		return
	}()

	for j := range 5 {
		jobs <- j
	}
	close(jobs)

	<-done
}
