package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println("Hello World!!")
		fmt.Println(texto)
	}("Walker")
}
