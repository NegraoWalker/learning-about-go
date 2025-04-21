package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Altura    float32
}

type Estudante struct {
	Pessoa    //"Herança" da struct Pessoa
	Curso     string
	Faculdade string
}

func main() {

	estudante1 := Estudante{
		Pessoa: Pessoa{ // Inicialização explícita do campo Pessoa
			Nome:      "Bruno",
			Sobrenome: "Santos",
			Idade:     26,
			Altura:    1.75,
		},
		Curso:     "Análise e Desenvolvimento de Sistemas",
		Faculdade: "Pitágoras",
	}

	fmt.Println("Estudante 1:", estudante1)

	var estudante2 Estudante
	estudante2.Nome = "Gabriel"
	estudante2.Sobrenome = "Felca"
	estudante2.Idade = 27
	estudante2.Altura = 1.80
	estudante2.Curso = "Dev"
	estudante2.Faculdade = "Uel"

	fmt.Println("Estudante 2: ", estudante2)

}
