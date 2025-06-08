package main

import "errors"

func main() {
	//Números: tipos inteiros:
	var inteiro1 int8 = 8
	var inteiro2 int16 = 16
	var inteiro3 int32 = 32
	var inteiro4 int64 = 64 //É equivalente ao int
	var inteiro5 uint8 = 8  //Sem sinal - Somente valores positivos

	//Números: tipos com ponto flutuante
	var float1 float32 = 32.67
	var float2 float64 = 64.67

	//Cadeia de caracteres: string
	var seqCaracteres string = "Texto usado para criar uma variável do tipo de dados string"

	//booleano:
	var ehVerdade bool = true

	//error:
	var erro error = errors.New("Erro interno!")

}
