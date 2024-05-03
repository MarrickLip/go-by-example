package main

import "fmt"

// function which returns an anonymous function. This forms a closure, which has state
func intSeq() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func main() {
	nextInt := intSeq()
	nextInt() // 1
	nextInt() // 2
	nextInt() // 3

	nextInt = intSeq()
	nextInt() // 4
}
