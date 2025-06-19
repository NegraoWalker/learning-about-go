package main

import "fmt"

func main() {
	//loop for tradicional
	for i := 0; i <= 5; i++ {
		fmt.Println("i:", i)
	}

	//loop for com um while
	j := 5
	for j >= 0 {
		fmt.Println("j:", j)
		j--
	}

	//loop for infinito
	input := ""
	for {
		fmt.Println("Digite 'sair' para encerrar o programa:")
		fmt.Scanln(&input)
		if input == "sair" {
			break
		}
	}

	//for range
	frutas := []string{"Maçã", "Banana", "Laranja"}
	for i, fruta := range frutas {
		fmt.Printf("Índice: %d, Fruta: %s\n", i, fruta)
	}
}
