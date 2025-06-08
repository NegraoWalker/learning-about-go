package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("walkereletrica@gmail.com") //Checa se o email o formato do email está correto a partir de uma regex
	fmt.Println(erro)
}
