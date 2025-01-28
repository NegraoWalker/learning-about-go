package main

import (
	"bytes"
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
	c := Cachorro{"Pipoca", "Dálmata", 3}
	fmt.Println(c)

	cEmJSON, erro := json.Marshal(c) //Conversão de struct para JSON
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cEmJSON)                  //Gera um slice de bytes
	fmt.Println(bytes.NewBuffer(cEmJSON)) //JSON que conseguimos ver
}
