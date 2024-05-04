package main

import "os"

func main() {
	_, err := os.Create("46-panic/test.txt")
	if err != nil {
		panic(err)
	}
}
