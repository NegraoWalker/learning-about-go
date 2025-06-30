package main

import "fmt"

//Função com retorno nomeado:
func calcular(numero1, numero2 int) (adicao, multiplicacao int) {
	adicao = numero1 + numero2
	multiplicacao = numero1 * numero2
	return
}

//Função com parâmetros variádicos:
func somar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Função recursiva:

func main() {
	ad, mult := calcular(2, 2)
	fmt.Println("Adição:", ad)
	fmt.Println("Multiplicação:", mult)

	fmt.Println("Soma:", somar(1, 1, 1, 1))

	//Função anônima:
	func(texto string) {
		fmt.Println(texto)
	}("Olá meu nome é Walker")

}
