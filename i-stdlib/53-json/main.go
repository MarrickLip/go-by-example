package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"currentPageIndex"`
	Fruits []string `json:"fruits"`
}

func main() {
	// encode primitives
	encodedBoolean, _ := json.Marshal(true)
	fmt.Printf("encodedBoolean: %v\n", string(encodedBoolean))

	encodedSlice, _ := json.Marshal([3]int{1, 2, 3})
	fmt.Printf("encodedSlice: %v\n", string(encodedSlice))

	// encode a struct
	res1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	encodedRes1, _ := json.Marshal(res1)
	fmt.Printf("encodedRes1: %v\n", string(encodedRes1))

	// enocode a struct with tags
	res2 := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	encodedRes2, _ := json.Marshal(res2)
	fmt.Printf("encodedRes2: %v\n", string(encodedRes2))

	rawResponse := `{"currentPageIndex": 12, "fruits": ["banana", "papaya"]}`
	res := response2{}
	err := json.Unmarshal([]byte(rawResponse), &res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %+v\n", res)
}
