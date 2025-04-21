package main

import "fmt"

type customString string //Definindo um alias para um tipo de dado personalizado

func (text customString) customLog() {
	fmt.Println(text)
}

func main() {
	var name customString
	name = "Walker"
	// fmt.Print(name)
	name.customLog()
}
