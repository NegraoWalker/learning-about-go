package main

import (
	"log"

	"github.com/NegraoWalker/myprogram/helpers"
)

func main() {

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
