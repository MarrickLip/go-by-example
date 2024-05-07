package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("p([a-z]+)ch", "peach")
	if err != nil {
		panic("MatchString failed")
	}
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)")
	matches := r.FindAllString("papaya and peach are pretty great fruits", -1)
	fmt.Printf("matches: %v\n", matches)
}
