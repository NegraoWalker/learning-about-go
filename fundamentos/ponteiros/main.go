package main

import "fmt"

func main() {
	var numero1 int = 12
	var p1 *int = &numero1 //Primeira forma de sintaxe de declaração de um ponteiro

	fmt.Println(p1)  //Imprime o endereço do local de memória da variável numero1
	fmt.Println(*p1) //Imprime o valor armazenado no local de memória da variável numero1

	var numero2 float64 = 6.56
	p2 := &numero2 //Segunda forma de sintaxe de declaração de um ponteiro

	fmt.Println(p2)  //Imprime o endereço do local de memória da variável numero2
	fmt.Println(*p2) //Imprime o valor armazenado no local de memória da variável numero2

	*p2 = 3.54           //Alterando o valor armazenado dentro da variável numero2
	fmt.Println(numero2) //Imprime o novo valor da variável numero2

}
