package main

import "fmt"

func main() {
	// slices are more flexible than arrays
	s := []string{"hello"}
	s = append(s, "world")
	fmt.Println(s, len(s))

	// slices have a capacity, which by default is the length
	fmt.Println(len(s), cap(s))

	// make allows slices to be initalised with a specific capacity
	t := make([]string, 5)
	fmt.Println(t, len(t), cap(t))

	// slice syntax ([high:low]) extracts a portion of a slice
	u := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(u[1:4])
}
