package main

import "fmt"

func main() {
	//Declaração de um slice:
	//Forma 1: A partir de um array
	arrayNotas := [5]float64{6.8, 7.4, 1.2, 10, 6}
	sliceNotas := arrayNotas[1:5]
	fmt.Println("Slice de notas:", sliceNotas)
	//Forma 2: Passando dados
	sliceNumerosInteiros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice de números inteiros", sliceNumerosInteiros)
	//Forma 3: Com make
	sliceNomes := make([]string, 3, 5)
	fmt.Println("Slice de nomes[Sem passar dados]:", sliceNomes)
	sliceNomes[0] = "Walker"
	sliceNomes[1] = "João"
	sliceNomes[2] = "Luis"
	fmt.Println("Slice de nomes:", sliceNomes)
	//Forma 4: Sem passar dados
	var sliceFrutas []string
	fmt.Println("Slice de frutas:", sliceFrutas)

	//Operações básicas:
	fmt.Println("Comprimento do slice de nomes:", len(sliceNomes))
	fmt.Println("Capacidade do slice de nomes:", cap(sliceNomes))
	sliceNomes[0] = "Luiza" //Alterando o dado de índice 0 do slice de nomes
	fmt.Println("Slice de nomes[alterado]:", sliceNomes)

	//Operações avançadas:
	sliceNomes = append(sliceNomes, "Lucas")
	sliceNomes = append(sliceNomes, "Silvio")
	sliceNomes = append(sliceNomes, "Isabela")
	fmt.Println("Slice de nomes[adicionado 3 novos dados]:", sliceNomes)
	fmt.Println("Comprimento do slice de nomes:", len(sliceNomes))
	fmt.Println("Capacidade do slice de nomes:", cap(sliceNomes))

}
