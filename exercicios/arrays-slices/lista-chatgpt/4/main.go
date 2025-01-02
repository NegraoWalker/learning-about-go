package main

import "fmt"

/*Crie um slice a partir de um array contendo [10, 20, 30, 40, 50]. O slice deve conter apenas os valores [20, 30, 40].
Imprima o slice e sua capacidade*/

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	slice := array[1:4]

	fmt.Println(slice)
}
