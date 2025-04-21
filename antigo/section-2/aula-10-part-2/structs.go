package main

import (
	"log"
	"time"
)

type User struct { //Declarando uma struct(estrutura)
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{FirstName: "Walker", LastName: "Negrão", PhoneNumber: "4333431097"}
	log.Println(user.FirstName)
	log.Println(user.LastName)
	log.Println(user.PhoneNumber)
	log.Println(user)
}
