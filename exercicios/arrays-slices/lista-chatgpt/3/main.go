package main

import "fmt"

/*Declare um array de tamanho 4 contendo os valores [10, 20, 30, 40]. Substitua o segundo elemento pelo valor 25 e imprima o array atualizado.*/

func main() {
	array := [4]int{10, 20, 30, 40}

	array[1] = 25

	fmt.Println(array)
}
