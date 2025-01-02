package main

import (
	"fmt"
	"slices"
)

/*Escreva uma função que receba um slice de inteiros e retorne o maior valor presente no slice.*/

func MaiorValor(slice []int) int {
	maiorValor := slices.Max(slice)
	return maiorValor
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(MaiorValor(slice))
}
