package main

import "log"

type Employee struct {
	Id       int
	Name     string
	Age      int
	Position string
}

func main() {

	var person Employee

	person.Id = 1
	person.Name = "Walker"
	person.Age = 31
	person.Position = "Doctor"

	log.Println(person)
	log.Println(person.Id)
	log.Println(person.Name)
	log.Println(person.Age)
	log.Println(person.Position)

	person2 := Employee{2, "Isabela", 30, "Human Resources"}
	log.Println(person2)

	person3 := Employee{
		Id:       3,
		Name:     "João",
		Age:      12,
		Position: "Teacher",
	}
	log.Println(person3)
}

/*
Struct: É um tipo de dado agregado que agrupa zero ou mais valores nomeados de tipos quaisquer como uma única entidade. Cada valor é chamado de campo.
*/
