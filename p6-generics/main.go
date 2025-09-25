package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	c := add(1, 2)
	fmt.Println("Result is: ", c)
	t := add("a", "b")
	fmt.Println("Result is: ", t)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
