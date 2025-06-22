package main

import "fmt"

func main() {
	//Sintaxe de declaração:
	//Forma 1:
	var notas [3]float64
	notas[0] = 7.8
	notas[1] = 9.2
	notas[2] = 6.5

	//Forma 2:
	frutas := [3]string{"Laranja", "Melancia", "Maçã"}

	//Forma 3:
	cores := [...]string{"azul", "verde", "amarelo"}

	//Acesso e modificação:
	notas[0] = 9.0
	fmt.Println("Array notas:", notas)

	fmt.Println("Array frutas - Primeiro dado armazenado:", frutas[0])

	cores[2] = "vermelho"
	fmt.Println("Array cores:", cores)

	//Tamanho:
	fmt.Println("Tamanho do array cores:", len(cores))

	//Sem inicialização:
	var idades [4]int
	fmt.Println("Array idades:", idades)
}
