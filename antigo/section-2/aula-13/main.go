package main

import "log"

func main() {
	// var isTrue bool
	// isTrue = false

	// if isTrue == true {
	// 	log.Println("isTrue is", isTrue)
	// } else {
	// 	log.Println("isTrue is", isTrue)
	// }

	// cat := "cat"

	// if cat == "cat" {
	// 	log.Println("Cat is cat!")
	// } else {
	// 	log.Println("Cat isn't cat!")
	// }

	// myNum := 100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println("Condition 1")
	// } else if myNum < 100 && isTrue {
	// 	log.Println("Condition 2")
	// } else if myNum == 101 || isTrue {
	// 	log.Println("Condition 3")
	// } else {
	// 	log.Println("Condition 4")
	// }

	myVar := ""

	switch myVar {
	case "cat":
		log.Println("Case 1")
	case "dog":
		log.Println("Case 2")
	case "fish":
		log.Println("Case 3")
	default:
		log.Println("Case dafault")
	}

}
