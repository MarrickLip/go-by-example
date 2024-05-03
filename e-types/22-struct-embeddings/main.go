package main

type animal struct {
	species string
}

func (a animal) isCat() bool {
	return a.species == "cat"
}

// structs can be embedded in other structs
type pet struct {
	animal
	name string
}

func main() {
	// base struct is initialised explicitly
	p := pet{
		animal: animal{species: "cat"},
		name:   "gatsby",
	}

	// fields for the embedded struct can be accessed directly
	println(p.species, p.name)

	// methods for the embedded struct can be accessed directly
	println(p.isCat())
}
