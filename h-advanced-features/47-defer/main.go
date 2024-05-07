package main

import (
	"fmt"
	"os"
)

func createFile(p string) (*os.File, error) {
	fmt.Println("createFile")
	f, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	return f, nil
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
	f, err := createFile("47-defer/test.txt")
	defer closeFile(f)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writeFile(f)
}
