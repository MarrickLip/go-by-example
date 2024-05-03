package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(n int) (int, error) {
	if n == 0 {
		return -1, &argError{n, "can't work with 0"}
	}
	fmt.Println(1 / n)
	return 1 / n, nil
}

func main() {
	_, err := f(0)

	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println("argError occured")
	}
}
