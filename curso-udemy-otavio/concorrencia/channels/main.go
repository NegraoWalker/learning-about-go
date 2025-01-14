package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escreverNaTela("Hello World!", canal)
	// for {
	// 	mensagem, canalAberto := <-canal
	// 	if !canalAberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa!")

}

func escreverNaTela(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
