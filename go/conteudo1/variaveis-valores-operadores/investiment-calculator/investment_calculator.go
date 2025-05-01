package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount int64 = 1000
	var expectedReturnRate float64 = 5.5
	var years int64 = 10

	var futureValue float64 = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Print(futureValue)
}
