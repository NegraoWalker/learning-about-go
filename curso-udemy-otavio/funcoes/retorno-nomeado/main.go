package main

import "fmt"

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma)
	fmt.Println(subtracao)
}

func calculosMatematicos(numero1, numero2 int) (soma, subtracao int) {
	soma = numero1 + numero2
	subtracao = numero1 - numero2
	return
}
