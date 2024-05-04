package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
)

func sortRandomArray() {
	nums := [1000]int{}
	for i := range nums {
		nums[i] = rand.Intn(1000000000)
	}

	slices.Sort(nums[:])
}

func main() {
	strs := []string{"b", "c", "a"}
	slices.Sort(strs)
	fmt.Println(strs)

	nums := []int{2, 3, 1}
	slices.Sort(nums)
	fmt.Println(nums)

	println("nums is sorted?", slices.IsSorted(nums))

	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sortRandomArray()
		}()
	}

	wg.Wait()
	fmt.Println("finished stress test")
}
