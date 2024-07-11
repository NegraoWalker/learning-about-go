package main

import "log"

type User struct { //Definindo uma struct para usar como um tipo de dado no map
	FirstName string
	LastName  string
	Age       int
}

func main() {
	myMap := make(map[string]string) //Sintaxe de criação de um map

	myMap["dog"] = "Scooby"
	myMap["other-dog"] = "Samson"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap["dog"] = "fido" //Modificando o valor armazenado para a chave dog
	log.Println(myMap["dog"])

	myMap2 := make(map[string]int)

	myMap2["first"] = 1
	myMap2["second"] = 2

	log.Println(myMap2["first"])
	log.Println(myMap2["second"])

	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Walker",
		LastName:  "Negrão",
		Age:       31,
	}

	myMap3["me"] = me

	log.Println(myMap3["me"].FirstName)

}
