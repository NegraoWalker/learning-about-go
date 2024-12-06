package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	error := checkmail.ValidateFormat("example@example.com")
	fmt.Println(error)
}
