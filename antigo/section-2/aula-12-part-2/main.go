package main

import "log"

func main() {

	var mySlice1 []string //1°Forma de declaração

	mySlice1 = append(mySlice1, "Walker")
	mySlice1 = append(mySlice1, "Esteves")
	mySlice1 = append(mySlice1, "Negrão")
	log.Println(mySlice1)

	var mySlice2 []int //1°Forma de declaração

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)
	log.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //2°Forma de declaração
	log.Println(numbers)
	log.Println(numbers[0:2])

}
