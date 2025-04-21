/*
	* Validar a entrada de dados:
		** Mostrar a mensagem de erro e interromper o fluxo de execução: Números negativos e 0
	* Gravar o valor calculado em um arquivo de texto
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, error1 := getUserInput("Revenue: ")
	expenses, error2 := getUserInput("Expenses: ")
	taxRate, error3 := getUserInput("Tax Rate: ")

	if error1 != nil || error2 != nil || error3 != nil {
		fmt.Print("ERROR: ", error1)
		fmt.Println()
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeToFile(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0.0, errors.New("value must be a positive number")
	}
	return userInput, nil
}

func storeToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nPROFIT: %.2f\nRATIO: %.2f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
