/*
Crie uma função swap que receba dois ponteiros para inteiros e troque os valores apontados por eles.
*/
package main

import "fmt"

func main() {
	numero1 := 1
	numero2 := 2
	increment(&numero1, &numero2)
	fmt.Println(numero1)
	fmt.Println(numero2)
}

func increment(p1, p2 *int) {
	aux := *p1
	*p1 = *p2
	*p2 = aux

	// *p1, *p2 = *p2, *p1 //Específico da linguagem Go
}
