package main

func main() {
	i := 13
	if i%2 == 0 {
		println("i is even")
	} else {
		println("i is odd")
	}

	// statements can precede conditionals, they are available inside the block
	if j := 13; j%2 == 0 {
		println("j is even")
		println("j is", j)
	} else {
		println("j is odd")
		println("j is", j)
	}

	// won't work:
	// println("j is", j)
}
