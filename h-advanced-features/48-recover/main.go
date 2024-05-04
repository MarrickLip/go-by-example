package main

import (
	"fmt"
	"math/rand"
)

func mightPanic() {
	if n := rand.Intn(10); n < 5 {
		fmt.Println(n, "is less than 5. Panicing!!")
		panic("number was too small")
	} else {
		fmt.Println(n, "is greater than or equal to 5")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	mightPanic()
}
