/*
Crie uma struct Carro com os campos Modelo, Ano e Cor. Crie uma instância dessa struct e imprima os valores.
*/

package main

import "fmt"

type Carro struct {
	Modelo string
	Ano    uint
	Cor    string
}

func main() {
	carro := Carro{
		Modelo: "Corsa",
		Ano:    1994,
		Cor:    "Rosa",
	}

	fmt.Println(carro)
}
