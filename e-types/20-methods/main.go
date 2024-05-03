package main

import "fmt"

type rect struct {
	width, height int
}

// methods have a reciever e.g. (r *rect)
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) isSquare() bool {
	return r.width == r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.area(), r.isSquare())
}
