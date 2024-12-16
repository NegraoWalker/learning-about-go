/*
Escreva uma função que receba um ponteiro para um slice de inteiros e adicione um valor ao slice.
*/
package main

import "fmt"

func main() {
	slice := []int{1, 2}
	appendValue(&slice, 3)
	fmt.Println(slice)
}

func appendValue(slice *[]int, value int) {
	*slice = append(*slice, value)
}
