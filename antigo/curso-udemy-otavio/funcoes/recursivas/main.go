package main

import "fmt"

func main() {
	fmt.Println(fibonacci(3))
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//0,1, 1, 2, 3, 5, 8, 13, 21
