package main

import "fmt"

/*Declare um array de 3 strings com os valores "Go", "Python" e "Java". Use um for loop para imprimir cada elemento do array.*/

func main() {
	array := [3]string{"Go", "Python", "Java"}

	for i := 0; i < 3; i++ {
		fmt.Println(array[i])
	}
}
