package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices = append(prices, 7.66)
	fmt.Println(prices)
}

//Forma 1:
// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00} //Declaração de um array de 4 posições float64
// 	fmt.Println("Array completo:", prices)
// 	fmt.Println("Posição 2 do array:", prices[2]) //Acessando elementos do array
// 	prices[0] = 11.99                             //Alterando elementos do array a partir do indice
// 	fmt.Println("Posição 0 do array", prices[0])

// 	featuredPrices := prices[1:3] //Forma 1 de declaração de slice
// 	fmt.Println("Elementos 1 e 2 do array:", featuredPrices)
// 	featuredPrices = prices[:3]
// 	fmt.Println("Elementos 0, 1 e 2 do array:", featuredPrices)
// 	featuredPrices = prices[1:]
// 	fmt.Println("Elementos 1, 2 e 3 do array:", featuredPrices)

// 	featuredPrices = prices[:3]
// 	featuredPrices[0] = 77.77 //Alterando o valor do elemento 0 do slice altera o array original tb

// 	fmt.Println(featuredPrices)
// 	fmt.Println(prices)

// }
