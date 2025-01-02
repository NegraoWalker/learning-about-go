package main

import "fmt"

func main() {
	// defer funcao1() //A palavra chave defer adia a execução da função 1. Assim, a função 2 será executada primeiro que a função 1
	// funcao2()
	fmt.Println(alunoEstaAprovado(7, 7.8))
}

func funcao1() {
	fmt.Println("Execução da função 1")
}
func funcao2() {
	fmt.Println("Execução da função 2")
}

func alunoEstaAprovado(nota1, nota2 float32) bool {
	defer fmt.Println("Media calculada")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado!")
	media := nota1 + nota2
	if media >= 7 {

		return true
	}
	return false
}
