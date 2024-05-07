package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("46-panic/test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Printf("f: %v\n", f)
}
