package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     uint8
}

type Usuario struct { //Declaração de uma struct
	Nome     string
	Idade    uint8
	Endereco Endereco //struct aninhada
}

func main() {
	var usuario1 Usuario //Forma 1
	fmt.Println("struct Usuario: ", usuario1)

	usuario1.Nome = "Walker"
	usuario1.Idade = 31
	usuario1.Endereco.Logradouro = "Rua Ernestina Duque Estrada"
	usuario1.Endereco.Numero = 180
	fmt.Println("struct Usuario - usuario1: ", usuario1)

	usuario2 := Usuario{ //Forma 2
		Nome:  "João",
		Idade: 22,
		Endereco: Endereco{
			Logradouro: "Flor dos Alpes",
			Numero:     100,
		},
	}
	fmt.Println("struct Usuario - usuario2: ", usuario2)
}
