package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct { //Declarando uma struct para receber os dados da resposta da API externa no formato de JSON
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := //Meu programa recebeu a resposta de uma API externa no formato JSON igual a esse exemplo
		`[
			{
				"first_name": "Clark",
				"last_name": "Kent",
				"hair_color": "black",
				"has_dog": true
			},
			{
				"first_name": "Bruce",
				"last_name": "Wayne",
				"hair_color": "black",
				"has_dog": false
			}
		]`

	var unmarshalled []Person //Declaração de uma slice

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled) //unmarshalled: [{Clark Kent black true} {Bruce Wayne black false}]

	//write json from a struct:
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
	/*
		[
			{
				"first_name": "Wally",
				"last_name": "West",
				"hair_color": "red",
				"has_dog": false
			},
			{
				"first_name": "Diana",
				"last_name": "Prince",
				"hair_color": "black",
				"has_dog": false
			}
		]
	*/
}
