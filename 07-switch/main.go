package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 0:
		fmt.Println("i is zero")
	case 1:
		fmt.Println("i is one")
	case 2:
		fmt.Println("i is two")
	default:
		fmt.Println("i is not zero, one, or two")
	}

	// switches can use conditionals instead of inputting a value
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	}

	// type switches compare the type of an interface, not values
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know the type")
		}
	}

	whatAmI(123)
	whatAmI("hello")

}
