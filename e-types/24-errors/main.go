package main

// in Go it's idiomatic to comunicate errors via a separate return value

import (
	"errors"
	"fmt"
)

func f(n int) (int, error) {
	if n == 3 {
		return -1, errors.New("can't work with 3")
	} else {
		return n, nil
	}
}

func main() {
	for i := range 10 {
		if res, err := f(i); err != nil {
			fmt.Println("f failed:", err)
		} else {
			fmt.Println("f worked:", res)
		}
	}
}
