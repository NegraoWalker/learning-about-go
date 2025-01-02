package main

import "fmt"

/*Crie dois slices: um contendo [1, 2, 3] e outro com tamanho suficiente para armazenar esses valores. Use a função copy
para copiar os valores do primeiro slice para o segundo e imprima ambos os slices.*/

func main() {
	slice1 := make([]int, 0)
	slice1 = append(slice1, 1, 2, 3)

	fmt.Println("slice 1:", slice1)
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	fmt.Println("slice 2:", slice2)

}
