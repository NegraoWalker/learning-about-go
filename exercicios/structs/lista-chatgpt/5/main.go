/*
Crie uma struct ContaBancaria com os campos Titular e Saldo. Crie uma função Depositar que recebe um ponteiro para a
struct e modifica o saldo.
*/

package main

import "fmt"

type ContaBancaria struct {
	Titular string
	Saldo   float64
}

func Depositar(ponteiro *ContaBancaria, valor float64) {
	ponteiro.Saldo += valor
}

func main() {
	contaBancaria := ContaBancaria{
		Titular: "Walker",
		Saldo:   0,
	}

	fmt.Println(contaBancaria)

	Depositar(&contaBancaria, 100)
	fmt.Println(contaBancaria)
}
