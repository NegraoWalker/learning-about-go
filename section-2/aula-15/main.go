package main

import "fmt"

type Animal interface { //Definindo uma interface
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jack",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(animal Animal) {
	fmt.Println("This animal says", animal.Says(), "and has", animal.NumberOfLegs(), "legs")
}

func (dog *Dog) Says() string {
	return "Woof"
}

func (dog *Dog) NumberOfLegs() int {
	return 4
}

func (gorilla *Gorilla) Says() string {
	return "Ugh"
}

func (gorilla *Gorilla) NumberOfLegs() int {
	return 2
}
