package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5 //constante global

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	f, rfv := calculateFutureValues(investmentAmount, expectedReturnRate, years) //funcão retorna dois valores

	fmt.Printf("Future Value: %.2f\n", f)
	fmt.Printf("Future Real Value: %v", rfv)
}

// Função:
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
