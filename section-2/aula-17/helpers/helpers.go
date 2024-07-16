package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(scale int) int {
	rand.NewSource(time.Now().UnixNano())
	value := rand.Intn(scale) //Função do package math que gera números aleatórios inteiros
	return value
}
