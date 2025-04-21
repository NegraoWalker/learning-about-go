package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0.00, errors.New("Failed to read file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0.0, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	var choice int
	var someCondition bool = true

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		// panic("Can't continue, sorry!!") //função para interromper o fluxo de execução caso ocorra algum erro
		fmt.Println("-----------------------")
		return //return interrompe o fluxo de execução
	}

	fmt.Println("Welcome to Go Bank!")
	for someCondition {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: US$", accountBalance)
		case 2:
			fmt.Print("How much do you want to deposit (US$): ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance updated, new amount: US$ %.3f", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("How much do you want to withdrawl (US$): ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Printf("Balance updated, new amount: US$ %.3f", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for choosing our bank!")
			someCondition = false
		default:
			fmt.Println("Invalid chosen option!")
		}
	}

}
