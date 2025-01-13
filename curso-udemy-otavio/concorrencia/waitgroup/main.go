package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escreverNaTela("Hello World!")
		waitGroup.Done()
	}()

	go func() {
		escreverNaTela("Programando em Go!!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escreverNaTela(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
