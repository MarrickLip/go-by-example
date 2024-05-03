package main

import "fmt"

func main() {
	// initialise an array which holds exactly 5 ints
	// by default it is zero-valued i.e. [0, 0, 0, 0, 0]
	var a [5]int
	fmt.Println(a)

	// len(x) gets the length of the array
	fmt.Println(a, "has", len(a), "elements")

	// initalise using {}; [...] tells the compiler to calculate how many enties
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// multi-dimensional arrays are supported
	var c [2][3]int
	for i := range 2 {
		for j := range 3 {
			c[i][j] = i + j
		}
	}

	fmt.Println(c)
}
