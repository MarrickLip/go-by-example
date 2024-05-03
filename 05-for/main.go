package main

import "fmt"

func main() {

	// the most basic for loop is essentially a while loop
	i := 1
	for i <= 3 {
		fmt.Println("basic", i)
		i += 1
	}

	// the variable can be intialised, compared, and incremented in one
	for j := 0; j < 3; j++ {
		fmt.Println("classic", j)
	}

	// range(n) increments from 1 to n
	for i := range 3 {
		fmt.Println("range", i)
	}

	// loops without a condition will continue indefinitely
	for {
		fmt.Println("loop")
		break
	}
}
