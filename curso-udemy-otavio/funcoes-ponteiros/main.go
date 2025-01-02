package main

import "fmt"

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero) //Faz uma cópia do valor da váriavel para a função utilizar
	fmt.Println(numeroInvertido)
	fmt.Println(numero) //Valor na váriavel permanece inalterado
	fmt.Println("======================================================================================================")
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) //Passando o endereço de memória da váriavel
	fmt.Println(novoNumero)
}

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}
