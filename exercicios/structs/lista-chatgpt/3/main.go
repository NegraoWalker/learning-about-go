/*
Dada uma struct Pessoa com os campos Nome e Idade, crie uma instância dessa struct, modifique o valor de Idade e
imprima os valores antes e depois da mudança.
*/

package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade uint
}

func main() {
	pessoa := Pessoa{
		Nome:  "Walker",
		Idade: 32,
	}

	fmt.Println(pessoa)

	pessoa.Idade = 31

	fmt.Println(pessoa)
}
