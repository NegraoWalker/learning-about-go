package main

import "fmt"

func soma(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func calcular(numero1, numero2 int64) (int64, int64) {
	divisao := numero1 / numero2
	multiplicacao := numero1 * numero2
	return divisao, multiplicacao
}

func main() {
	fmt.Println(soma(1, 5)) //1

	subtrai := func(numero1 int16, numero2 int16) int16 {
		return numero1 - numero2
	}
	fmt.Println(subtrai(4, 2)) //2

	resultadoDiv, resultadoMult := calcular(16, 4) //3
	fmt.Printf("Divisão: %d\n", resultadoDiv)
	fmt.Printf("Multiplicação: %d\n", resultadoMult)

	resultadoDiv2, _ := calcular(40, 10) //Não quero receber o segundo resultado
	fmt.Printf("Divisão: %d\n", resultadoDiv2)
}
