package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero uint
}

//Declaração da struct
type Usuario struct {
	Nome     string
	Idade    uint
	Endereco //Conceito de "herança" para Go - Utilizar outra struct dentro de outra por meio da composição
}

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint
	Altura    float64
}

//Conceito de "herança" para Go
type Estudante struct {
	Pessoa
	Curso     string
	Faculdade string
}

func main() {
	//Instanciação da struct Usuario
	usuario1 := Usuario{
		Nome:     "Walker",
		Idade:    32,
		Endereco: Endereco{"Ernestina Duque Estrada", 180},
	}

	fmt.Println("Nome do usuário 1:", usuario1.Nome)
	fmt.Println("Idade do usuário 1:", usuario1.Idade)
	fmt.Println("Endereço do usuário 1:", usuario1.Endereco)

	estudante1 := Estudante{
		Pessoa: Pessoa{
			Nome:      "João",
			Sobrenome: "Silva",
			Idade:     45,
			Altura:    1.75,
		},
		Curso:     "Engenharia Mecânica",
		Faculdade: "UEL",
	}

	fmt.Println("Nome do estudante 1:", estudante1.Nome)
	fmt.Println("Curso do estudante 1:", estudante1.Curso)
	fmt.Println("Faculdade do estudante 1:", estudante1.Faculdade)
}
