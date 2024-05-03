package main

func sum(nums ...int) (int, int) {
	total := 0
	n := 0
	for _, num := range nums {
		total += num
		n += 1
	}

	return total, n
}

func main() {
	total, n := sum(1, 2, 3, 4, 5)
	println(total, n)
}
