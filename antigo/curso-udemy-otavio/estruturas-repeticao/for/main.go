package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Iteração i:", i)
	}

	fmt.Println("=========================================================================")
	j := 0
	for j < 10 {
		fmt.Println("Iteração j:", j)
		j++
	}

	fmt.Println("=========================================================================")
	nomes := [4]string{"Isabela", "Walker", "Emerson", "Beatriz"}
	for indice, nomes := range nomes {
		fmt.Println(indice, nomes)
	}

}
