package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("createFile")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writeFile")
	fmt.Fprintln(f, "hello world")
}

func closeFile(f *os.File) {
	fmt.Println("closeFile")
	f.Close()
}

func main() {
	f := createFile("47-defer/test.txt")
	defer closeFile(f)

	writeFile(f)
}
