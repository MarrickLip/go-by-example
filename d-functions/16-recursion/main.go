package main

import (
	"fmt"
)

func fact(n int) int {
	if n <= 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	fmt.Println(fact(25))

	// closures can recurse, but they must be declared explicitly with a typed var first
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}

	fmt.Println(fib(25))
}
