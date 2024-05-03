package main

import "fmt"

type person struct {
	name string
	age  int
}

// it's idiomatic to construct structs in a function which returns a pointer
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 50
	return &p
}

func main() {
	a := person{"Arthur Dent", 42}
	b := newPerson("Ford Prefect")

	fmt.Println(a)
	fmt.Println(b)

	// structs are mutable
	a.age += 1
	fmt.Println("Happy birthday Arthur! You're now", a.age)
}
