/*
Implemente uma função chamada increment que receba um ponteiro para um número inteiro e incremente o valor apontado por ele.
*/
package main

import "fmt"

func main() {
	numero := 0
	increment(&numero) //Passando o endereço da variável numero
	fmt.Println(numero)
}

func increment(p *int) {
	*p++
}
