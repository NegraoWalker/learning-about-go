package main

import "fmt"

func main() {
	age := 32
	// var agePointer *int //Declaração de um ponteiro
	// agePointer = age
	agePointer := &age //Passando o endereço de memória da variável age

	fmt.Println("Valor armazenado na variável age: ", age)
	fmt.Println("Endereço de memória da variável age: ", agePointer)
	getAdultYears(agePointer)
	fmt.Println("Retorno da função getAdultYears: ", age)
	// fmt.Println("Retorno da função getAdultYears: ", getAdultYears(agePointer)) //ou &age
	fmt.Println("Valor armazenado na variável age, usando o operador * no ponteiro: ", *agePointer)
}

func getAdultYears(age *int) { //O parametro age da função é uma cópia do valor da variável age. Por ser uma cópia são endereços de memória diferentes
	// return *age - 10
	*age = *age - 10
} //age *int é a declaração de um ponteiro
