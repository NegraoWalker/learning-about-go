package main

import (
	"fmt"
	"time"
)

func main() {
	go escreverNaTela("Hello World!") //goroutine
	escreverNaTela("Programando em Go!!")
}

func escreverNaTela(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
