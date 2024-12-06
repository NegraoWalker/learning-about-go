package main

import "fmt"

func main() {
	var nome string = "Walker"            //Declaração de tipo explícita
	var idade int64 = 31                  //Declaração de tipo explícita
	endereco := "Ernestina Duque Estrada" //Declaração para inferência de tipo pelo compilador
	salario := 1569.344                   //Declaração para inferência de tipo pelo compilador

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(endereco)
	fmt.Println(salario)

}
