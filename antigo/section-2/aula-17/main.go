package main

import (
	"log"

	"github.com/NegraoWalker/mod-aula-17/helpers"
)

const numPool = 10

func CalculateValue(intChannel chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChannel <- randomNumber
}

func main() {
	intChannel := make(chan int) //Criação de um channel(canal), abrindo o canal
	defer close(intChannel)      //Fechando o canal

	go CalculateValue(intChannel) //Iniciando um goRoutine

	num := <-intChannel //Recebendo o contéudo transmitido pelo canal
	log.Println(num)
}
