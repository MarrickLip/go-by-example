package main

import "fmt"

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head *element[T]
	tail *element[T]
}

func (list *List[T]) Push(val T) {
	if list.tail == nil {
		list.head = &element[T]{val: val}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: val}
		list.tail = list.tail.next
	}
}

func (list *List[T]) PrintAll() {
	for e := list.head; e != nil; e = e.next {
		fmt.Print(e.val)
		fmt.Print(" ")
	}
}

func main() {
	list := List[int]{}

	list.Push(1)
	list.Push(2)
	list.Push(3)

	list.PrintAll()
}
