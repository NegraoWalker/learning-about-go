package main

import "fmt"

var n int

func main() {
	fmt.Println("Execução da função main()")
	fmt.Println("Valor de n:", n)
}

func init() {
	fmt.Println("Execução da função init()")
	n = 10
}
