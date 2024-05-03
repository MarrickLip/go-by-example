package main

import "fmt"

func main() {
	// create a map using make(map[key-type]val-type)
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["z"] = 26

	fmt.Println(m)

	// delete(map, key) deletes a key from map
	delete(m, "z")

	fmt.Println(m)

	// the second return value indicates whether the item exists
	// n.b if it doesn't, val will be the zero value

	i, ok := m["0"]
	fmt.Println(i, ok)

	// maps can be initialised inline

	n := map[string]string{"h": "hello", "w": "world"}
	println(n["h"], n["w"])

}
