package main

import (
	"fmt"
)

func main() {

	fmt.Println("Arrays:")

	var array1 [2]string //Declaração de um array de tamanho 2 (armazena até 2 dados)
	array1[0] = "Hello"
	array1[1] = "World"
	fmt.Println(array1[0], array1[1])
	fmt.Println(array1)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13} //Declaração de um array de tamanho 6(armazena até 6 dados)
	fmt.Println(numPrimos)
	fmt.Println(len(numPrimos))
	fmt.Println(numPrimos[0:5])
	fmt.Println(numPrimos[0:6])
	fmt.Println(numPrimos[1:3])

	fmt.Println("Slices:")

	slice1 := make([]string, 5)
	slice1[0] = "Hello"
	slice1[1] = "World"
	fmt.Println(slice1[0], slice1[1])
	fmt.Println(slice1)

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

}

/*
	Array e Slice são estruturas de dados homogêneas (armazenam somente dados do mesmo tipo)
	Array:
		Possue tamanho fixo, de zero ou mais elementos do mesmo tipo;
		Acessamos os valores com índice: a[0], a[1]...
		A função embutida len() retorna o tamanho do array;
		Por conta de ter tamanho fixo, não é tão usada. Só em casos específicos;
	Slice:
		Possue tamanho variável, de zero ou mais elementos do mesmo tipo;
		Acessamos os valores com índice: s[0], s[1]...
		A função embutida len() retorna o tamanho do slice;
		A função append() é usada para adicionar novos elementos


*/
