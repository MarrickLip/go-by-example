package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	p := point{1, 2}

	// %v prints just the value
	fmt.Printf("p: %v\n", p)

	// %+v inlcludes the field names in the struct
	fmt.Printf("p: %+v\n", p)

	// %T prints the type of the value
	fmt.Printf("p: %T\n", p)
}
