package main

import "fmt"

func main() {
	fmt.Println("=====================================================================================================================")
	var variavel1 int
	var ponteiro1 *int //Declaração do ponteiro
	fmt.Println("Valor variavel 1: ", variavel1)
	fmt.Println("Valor ponteiro 1: ", ponteiro1)
	fmt.Println("=====================================================================================================================")
	/*---------------------------------------------------------------------------------------------------------------------------------*/
	variavel1 = 99
	ponteiro1 = &variavel1
	fmt.Println("Valor variavel 1: ", variavel1)
	fmt.Println("Valor do endereço do ponteiro 1: ", ponteiro1) //Apontando para o endereço de memória da variavel 1
	fmt.Println("Valor ponteiro 1: ", *ponteiro1)               //Processo de desreferenciação, exibe o valor armazenado na variável 1 e não mais o endereço
	fmt.Println("=====================================================================================================================")
	/*---------------------------------------------------------------------------------------------------------------------------------*/

}
