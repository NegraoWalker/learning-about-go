package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { //Método da estrutura usuario
	fmt.Println("Dentro do método salvar")
}

type Contador struct {
	Valor int
}

// Método com receptor de ponteiro
func (c *Contador) Incrementar() {
	c.Valor++
}

func main() {
	usuario := usuario{"user1", 20}
	usuario.salvar()
	fmt.Println(usuario)

	c := Contador{Valor: 0}
	fmt.Println("Antes:", c.Valor)
	c.Incrementar()
	fmt.Println("Depois:", c.Valor)
}
