/*
Crie um programa que: Declare um array de inteiros [5]int. Use um ponteiro para alterar o valor do terceiro elemento do array para 50.Imprima o array atualizado.
*/
package main

import "fmt"

func main() {
	array := [5]int{0, 1, 2, 3, 4}
	var p *int = &array[2]

	fmt.Println(array)
	*p = 50 //Alterando o valor do índice 2 do array via ponteiro
	fmt.Println(array)

}
