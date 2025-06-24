package main

import "fmt"

func main() {
	//Declaração:
	m1 := map[string]int{
		"Laranja":  10,
		"Banana":   20,
		"Melancia": 2,
	}

	fmt.Println(m1)

	//Inserindo um novo valor:
	m1["Goiaba"] = 14
	//Atualizando um valor:
	m1["Laranja"] = 25
	//Acessando um valor:
	qtdLaranja := m1["Laranja"]
	fmt.Println(qtdLaranja)
	//Excluindo um dado:
	delete(m1, "Banana")
	fmt.Println(m1)
	//Tamanho do map:
	fmt.Println(len(m1))
	//Iterando sobre o map:
	for chave, valor := range m1 {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}
