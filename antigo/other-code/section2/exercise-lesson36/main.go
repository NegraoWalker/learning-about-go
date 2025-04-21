package main

import "fmt"

func main() {

	revenue, expenses, taxRate := getInput("Revenue: ", "Expenses: ", "Tax Rate: ")
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getInput(infoText1, infoText2, infoText3 string) (float64, float64, float64) {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print(infoText1)
	fmt.Scan(&revenue)
	fmt.Print(infoText2)
	fmt.Scan(&expenses)
	fmt.Print(infoText3)
	fmt.Scan(&taxRate)
	return revenue, expenses, taxRate
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
