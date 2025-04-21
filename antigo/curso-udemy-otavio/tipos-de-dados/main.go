package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Inteiros:
	var inteiro1 int = 1 //64 bits tb
	// var inteiro2 int8 = 2
	// var inteiro3 int16 = 3
	// var inteiro4 int32 = 4
	// var inteiro5 int64 = 5

	//Reais:
	var real1 float32 = 1.11
	// var real2 float64 = 2.22

	//Strings:
	var string1 string = "ABC#123"

	//Booleanos:
	var booleano1 bool = true

	//Os demais procurar na ducumentação oficial

	fmt.Println(reflect.TypeOf(inteiro1))
	fmt.Println(reflect.TypeOf(real1))
	fmt.Println(reflect.TypeOf(string1))
	fmt.Println(reflect.TypeOf(booleano1))

}
