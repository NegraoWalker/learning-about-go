package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) //& aponta para o endereço de memória da variável myString
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) { //*string indica que a função espera receber um ponteiro para string no parâmetro s
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue //*s aponta para o endereço na memória onde o valor está armazenado
}
