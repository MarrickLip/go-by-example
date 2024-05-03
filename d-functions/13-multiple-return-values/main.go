package main

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	println(a, b)

	// use _ to ignore values
	_, c := vals()
	println(c)
	// won't work: println(_)
}
