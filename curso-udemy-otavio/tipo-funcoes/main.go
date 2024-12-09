package main

import "fmt"

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultado := somar(2, 3)
	fmt.Println("Função somar:", resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 2)
	fmt.Println("Função calculos matemáticos: Soma:", resultadoSoma)
	fmt.Println("Função calculos matemáticos: Subtração:", resultadoSubtracao)
}
