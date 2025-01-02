package main

import "fmt"

func main() {
	numero1 := 9

	if numero1 > 10 {
		fmt.Println("Caiu na condição 1!")
	} else if numero1 == 10 {
		fmt.Println("Caiu na condição 2!")
	} else {
		fmt.Println("Caiu na condição 3!")
	}

	dia := "segunda"

	switch dia {
	case "segunda":
		fmt.Println("Início da semana!")
	case "sexta":
		fmt.Println("Quase fim de semana!")
	case "sábado", "domingo":
		fmt.Println("Fim de semana!")
	default:
		fmt.Println("Dia comum.")
	}
}
