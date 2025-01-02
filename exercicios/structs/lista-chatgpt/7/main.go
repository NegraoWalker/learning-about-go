/*
Crie uma struct Retangulo com os campos Largura e Altura. Implemente um método Area que calcule a área do retângulo.
*/

package main

import "fmt"

type Retangulo struct {
	Largura float64
	Altura  float64
}

// Definição do método Area para a estrutura Retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func main() {
	retangulo := Retangulo{
		Largura: 10.62,
		Altura:  4.65,
	}

	fmt.Println(retangulo.Area())
}
