package main

import (
	"log"
)

var s = "seven" //Escopo global

func main() {
	log.Println("s is", s)
	var s2 = "six" //Escopo local
	log.Println("s2 is", s2)
	log.Println(saySomething("xxx"))
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s) //Valor de s disponível para uso aqui
	return s3, "World!"
}
