package main

import "fmt"

// Declaração de constantes individuais
const Pi = 3.14159      // constante numérica
const Version = "1.0.0" // constante string
const Monday = 1        // dias da semana, valores explícitos
const Tuesday = 2
const Wednesday = 3

func main() {
	// Declarações de variáveis com var e inicializador
	var x int = 10       // tipo explícito, inicializado em 10
	var y float64 = 20.5 // tipo explícito, inicializado em 20.5
	var z = x + int(y)   // tipo inferido, com conversão de y para int

	name := "Walker" // inferido como string
	greeting := fmt.Sprintf("Olá, %s!", name)

	// Uso das constantes e variáveis
	fmt.Println("Pi:", Pi)
	fmt.Println("Versão:", Version)
	fmt.Println("Dias:", Monday, Tuesday, Wednesday)
	fmt.Println("Valores:", x, y, z)
	fmt.Println(greeting)
}
