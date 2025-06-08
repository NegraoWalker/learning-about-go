package main

import "fmt"

func main() {
	var nome string = "Walker" //Informando o tipo
	fmt.Println(nome)
	var idade int64 = 32 //Informando o tipo
	fmt.Println(idade)

	sobrenome := "Negrão" //Inferência de tipo
	fmt.Println(sobrenome)
	salario := 3054.89 //Inferência de tipo
	fmt.Println(salario)

	const ANO_LETIVO = 2025
	fmt.Println(ANO_LETIVO)

}
