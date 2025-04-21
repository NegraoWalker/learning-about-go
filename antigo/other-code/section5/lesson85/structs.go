package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

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

	appUser, error := user.New(userFirstName, userLastName, userBirthdate) //Usando a função construtora para iniciar a struct user
	appAdmin := user.NewAdmin("antoniotest@example.com", "admin123")

	if error != nil {
		fmt.Println("error: ", error)
		return
	}

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	appAdmin.OutputUserDetails()
	appAdmin.ClearUserName()
	appAdmin.OutputUserDetails()

	fmt.Println(appAdmin.Email)
	fmt.Println(appAdmin.Password)

}
