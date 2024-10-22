package main

import "fmt"

func main() {
	result := add(1, 2)

	fmt.Println("Result Add:", result)
}

func add[T int | float64 | string](a, b T) T { //Generics
	return a + b
}
