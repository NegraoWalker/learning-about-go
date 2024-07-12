package main

import "log"

func main() {
	// for i := 0; i < 10; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// for i, animal := range animals { //i -> índice animal -> valor armazenado em cada índice na slice animals
	// 	log.Println(i, animal)
	// }
	// for _, animal := range animals { //_ em go é usado para ignorar
	// 	log.Println(animal)
	// }

	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }

	// var firstLine string = "Once upon a midnight dreary"

	// for i, line := range firstLine {
	// 	log.Println(i, ":", line)
	// }

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 19})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 56})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 25})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}

}
