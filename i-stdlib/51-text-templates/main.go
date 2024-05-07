package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is: {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1.Execute(os.Stdout, "1232132")
}
