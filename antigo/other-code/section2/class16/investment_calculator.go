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

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Real Value: %v", futureRealValue)
}
