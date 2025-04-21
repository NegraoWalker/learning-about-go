package main

import "fmt"

func main() {
	usuario1 := map[string]string{ // Declaração de um map
		"nome":      "Pedro",
		"sobrenome": "Batista",
	}

	fmt.Println(usuario1)

	usuario2 := map[string]map[string]string{ // Nesse caso o valor do map é outro map
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"status": "Concluido",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
