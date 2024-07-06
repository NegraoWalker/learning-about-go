package main

import "fmt"

func main() {
	if 8%2 == 1 {
		fmt.Println("Condição if")
	} else if 8%2 == 0 {
		fmt.Println("Condição else if")
	} else {
		fmt.Println("Condição else")
	}
}
