package main

import "fmt"

/*Crie um slice com os valores [5, 10, 15, 20]. Em seguida, adicione o valor 25 e remova o segundo elemento do slice.
Imprima o slice final.*/

func main() {

	slice := []int{5, 10, 15, 20}

	slice = append(slice, 25) //Adicionando um novo elemento
	fmt.Println(slice)

	slice = append(slice[:1], slice[2:]...) // slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

	fmt.Println(slice)
}
