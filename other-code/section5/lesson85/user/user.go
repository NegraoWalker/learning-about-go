package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { //Aqui o nome pode começar com letra minuscula ou maiuscula. Caso escolha maiuscula essa struct pode ser importada para outros arquivos
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

func New(firstName, lastName, birthdate string) (*User, error) { //Função construtora que inicia uma struct user
	if firstName == "" || lastName == "" || birthdate == "" { //Adicionando um validação dentro da função construtora
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ANTONIO",
			LastName:  "TEST",
			Birthdate: "-----",
			CreatedAt: time.Now(),
		},
	}
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

func (u *User) OutputUserDetails() { //receiver
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
