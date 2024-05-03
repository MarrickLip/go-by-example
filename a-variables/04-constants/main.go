package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	// numeric constants don't have a type until given one, either by:
	//   - explicit assignment e..g int64(x), or
	//   - being passed into a function with that type e.g. math.Sin(x)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
