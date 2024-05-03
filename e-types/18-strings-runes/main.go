package main

import "fmt"

func main() {
	const s = "hello, world"

	for _, val := range s {
		fmt.Printf("%x ", val)
	}

	fmt.Println()

	// a rune is just an integer which represents a unicode code point
	rune := 'a'
	fmt.Println(rune)
}
