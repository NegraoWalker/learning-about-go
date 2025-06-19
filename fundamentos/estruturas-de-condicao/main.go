package main

import "fmt"

func main() {
	//if
	idade1 := 18
	if idade1 >= 18 {
		fmt.Println("Maior de idade!")
	}
	//if-else
	idade2 := 15
	if idade2 >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade!")
	}
	//if-else if-else
	idade3 := 13
	if idade3 >= 18 {
		fmt.Println("Adulto")
	} else if idade3 >= 13 || idade3 < 18 {
		fmt.Println("Adolescente")
	} else {
		fmt.Println("Criança")
	}
	//switch
	mes := 6
	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abriu")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	}
}
