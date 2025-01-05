package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 { //Função associada a struct retangulo
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 { //Função associada a struct circulo
	return math.Pi * c.raio * c.raio
}

func exibirArea(f forma) {
	fmt.Printf("Área: %.2f\n", f.area())
}

func main() {
	r := retangulo{altura: 2.89, largura: 2.5}
	c := circulo{raio: 3.562}

	exibirArea(r)
	// fmt.Println()
	exibirArea(c)

}
