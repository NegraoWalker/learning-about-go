package main

import "fmt"

/*Crie um slice vazio de inteiros usando a função make. Adicione os valores 5, 10 e 15 ao slice. Em seguida, imprima o
slice e seu comprimento.*/

func main() {
	slice := make([]int, 0)
	fmt.Println(slice)

	slice = append(slice, 5, 10, 15)
	fmt.Println(slice)
}
