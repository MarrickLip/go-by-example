package main

import "fmt"

// i is passed by pointer
func zeroptr(i *int) {
	fmt.Println("the memory address of i is", &i)
	*i = 0 // *i deferences i
}

func main() {
	i := 1
	zeroptr(&i)

	// i is mutated
	fmt.Println("i", i)
}
