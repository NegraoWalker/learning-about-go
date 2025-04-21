package main

import "fmt"

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := clousure()
	funcaoNova()
}

func clousure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
