package main

func main() {
	// range for a slice
	nums := []int{10, 23, 51, 42, 23}
	for i, num := range nums {
		println(i, num)
	}

	// range for a map
	kv := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kv {
		println(key, value)
	}
}
