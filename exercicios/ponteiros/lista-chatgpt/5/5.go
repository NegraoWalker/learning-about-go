/*
Crie uma struct Person com os campos Name e Age. Crie uma função que receba um ponteiro para Person e modifique o nome.
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func main() {
	person := Person{
		Name: "Bruno",
		Age:  26,
	}

	changeName(&person, "Lucas")
	fmt.Println(person)
}

func changeName(person *Person, newName string) {
	person.Name = newName
}
