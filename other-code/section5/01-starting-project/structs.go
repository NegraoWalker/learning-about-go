package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct { //Aqui o nome pode começar com letra minuscula ou maiuscula. Caso escolha maiuscula essa struct pode ser importada para outros arquivos
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error) { //Função construtora que inicia uma struct user
	if firstName == "" || lastName == "" || birthdate == "" { //Adicionando um validação dentro da função construtora
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	// var appUser user

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, error := newUser(userFirstName, userLastName, userBirthdate) //Usando a função construtora para iniciar a struct user

	if error != nil {
		fmt.Println("error: ", error)
		return
	}

	// outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

func (u *user) outputUserDetails() { //receiver
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}
