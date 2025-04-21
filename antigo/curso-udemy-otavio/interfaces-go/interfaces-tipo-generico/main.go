package main

import "fmt"

func main() {
	interfaceGenerica("String")
	interfaceGenerica(12)
	interfaceGenerica(9.883)
	interfaceGenerica(true)
}

func interfaceGenerica(inter interface{}) { //Recebe qualquer tipo de dado
	fmt.Println(inter)
}
