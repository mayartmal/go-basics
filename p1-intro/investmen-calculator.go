package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// var investmentAmount float64
	// var years float64
	// var expectedReturnRate float64

	investmentAmount := valueSetter("Investment amount: ")
	years := valueSetter("Years: ")
	expectedReturnRate := valueSetter("Expected Return Rate: ")

	futureValue, futureRealValue := caclcFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %.1f\n", futureValue)
	fmt.Printf("Future Real Value: %.2f\n", futureRealValue)

}

func valueSetter(askingString string) float64 {
	var userInput float64
	fmt.Printf("%v: ", askingString)
	fmt.Scan(&userInput)
	return userInput
}

func caclcFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}
