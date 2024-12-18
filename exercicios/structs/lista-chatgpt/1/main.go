/*
Crie uma struct chamada Livro, que tenha os seguintes campos: Titulo, Autor, AnoPublicacao e NumeroPaginas
*/

package main

import "fmt"

type Livro struct {
	Titulo        string
	Autor         string
	AnoPublicacao uint
	NumeroPaginas uint
}

func main() {
	var livro Livro
	livro.Titulo = "XXXX"
	livro.Autor = "XX"
	livro.AnoPublicacao = 2001
	livro.NumeroPaginas = 300

	fmt.Println(livro)
}
