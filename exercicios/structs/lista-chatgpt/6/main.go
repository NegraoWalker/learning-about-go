/*
Crie uma struct Turma com um campo Alunos que seja um array de strings. Imprima os alunos da turma.
*/

package main

import "fmt"

type Turma struct {
	Alunos [3]string
}

func main() {
	turma := Turma{
		Alunos: [3]string{"Walker", "Lucas", "Mônica"},
	}

	fmt.Println(turma)
}
