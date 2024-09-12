package main

import "fmt"

func main() {
	var choice int
	var accountBalance = 1000.00

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	fmt.Println("Your choice:", choice)

	if choice == 1 {
		fmt.Println("Your balance is: US$", accountBalance)
	} else if choice == 2 {
		fmt.Print("How much do you want to deposit (US$): ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("Balance updated, new amount: US$ %.3f", accountBalance)
	}
}
