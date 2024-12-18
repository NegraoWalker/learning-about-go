/*
Crie uma função ExibirPessoa que recebe uma struct Pessoa como parâmetro e imprime seu nome e idade.
*/

package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade uint
}

func ExibirPessoa(pessoa Pessoa) {
	fmt.Println("Nome: ", pessoa.Nome)
	fmt.Println("Idade: ", pessoa.Idade)
}

func main() {
	pessoa := Pessoa{
		Nome:  "Gustavo",
		Idade: 19,
	}

	ExibirPessoa(pessoa)
}
