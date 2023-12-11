package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureInvestmentValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureInvestmentRealValue := futureInvestmentValue * math.Pow(1+inflationRate/100, years)

	fmt.Println(futureInvestmentValue)
	fmt.Println(futureInvestmentRealValue)
}
