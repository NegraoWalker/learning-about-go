package main

import "fmt"

func main() {

	var variavelInicializada1 string = "String"
	var variavelInicializada2 int16 = 1250
	var inteira1, inteira2, inteira3 int8
	var booleana1, decimal1, string3 = true, 2.344, "três"

	var variavel3 int     //valor padrão 0
	var variavel4 float32 //valor padrão 0
	var variavel5 string  //valor padrão vazio
	var variavel6 bool    //valor padrão false

	nome := "Walker" //short variable declaration
	idade := 31
	salario := 2950.569
	masculino := true

	fmt.Println(variavelInicializada1)
	fmt.Println(variavelInicializada2)
	fmt.Println(inteira1)
	fmt.Println(inteira2)
	fmt.Println(inteira3)
	fmt.Println(booleana1)
	fmt.Println(decimal1)
	fmt.Println(string3)

	fmt.Println("===============================================================")

	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	fmt.Println("===============================================================")
	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(salario)
	fmt.Println(masculino)
}
