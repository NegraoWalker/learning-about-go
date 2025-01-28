package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{
		"nome": "Rex",
		"raca": "Dálmata",
		"idade": 3
	}`

	var c Cachorro

	erro := json.Unmarshal([]byte(cachorroEmJSON), &c) //Conversão do dado em JSON para Struct
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}
