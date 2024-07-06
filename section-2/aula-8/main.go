package main

import "fmt"

func main() {
	var whatToSay string
	whatToSay = "Goodbye, cruel world"

	var i int
	i = 10

	whatWasSaid := saySomething() //declaração com inferência de tipo (váriavel)
	whatWasSaid2, theOtherThingThatWasSaid := saySomething2()

	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)
	fmt.Println(saySomething())
	fmt.Println("The function returned: ", whatWasSaid)
	fmt.Println("The function returned 2: ", whatWasSaid2, theOtherThingThatWasSaid)
}

func saySomething() string {
	return "Something"
}

func saySomething2() (string, string) {
	return "Something", "else"
}
