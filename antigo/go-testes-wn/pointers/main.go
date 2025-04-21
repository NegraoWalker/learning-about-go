package main

import (
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	myVar := 10
	log.Println(&myVar) //Utilizando o operador & consigo acessar o valor do endereço de memória(A base é hexadecimal) que está alocado a variável myVar
	log.Println(myVar)

	var person Person
	log.Println("Endereço de person", &person) //Não aparece o endereço
	log.Println(&person.FirstName)             // 0xc00002c150
	log.Println(&person.LastName)              //0xc00002c160
	log.Println(&person.Age)                   //0xc00002c170

	fmt.Println("====================================POINTER=========================================")
	myVar2 := 12.678       //Variável
	var myPointer *float64 //Declarando um ponteiro para float32

	log.Println(&myVar2)    //Endereço de memória de myVar2
	log.Println(&myPointer) //Endereço de memória do ponteiro
	log.Println(myVar2)     //Valor armazenado em myVar2
	log.Println(myPointer)  //Valor armazenado no ponteiro

	myPointer = &myVar2
	log.Println("Valor armazenado na variável myVar2:", myVar2)                               //Valor armazenado em myVar2                                                                    //Apontando para o endereço da variável myVar2
	log.Println("Endereço de memória da variável myVar2:", &myVar2)                           //Endereço de memória de myVar2
	log.Println("Endereço de memória que o ponteiro aponta:", myPointer)                      //Endereço de memória que o meu ponteiro armazena ou aponta
	log.Println("Valor armazenado no endereço de memória que o ponteiro aponta:", *myPointer) //Valor armazenado no endereço de memória que o ponteiro aponta

	//Alterando o valor da variável myVar2 usando o ponteiro criado:
	*myPointer = 11.56
	log.Println("O novo valor da variável myVar2 usando ponteiro é:", myVar2)

}
