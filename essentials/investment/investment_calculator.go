package main

import (
	"fmt"
	"math"
)

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) float64 {
	return investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
}

func calculateFutureValueWithInflation(investmentAmount float64, expectedReturnRate float64, years float64, inflationRate float64) float64 {
	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years);
	return futureValue / math.Pow(1 + inflationRate /  100, years)
}

func main() {
	const inflationRate float64 = 2.5;

	var investmentAmount, years, expectedReturnRate float64;

	fmt.Print("Investment amount: ");
	fmt.Scan(&investmentAmount);

	fmt.Print("Years: ");
	fmt.Scan(&years);

	fmt.Print("expectedReturnRate: ");
	fmt.Scan(&expectedReturnRate);

	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years);
	futureRealValue := calculateFutureValueWithInflation(investmentAmount, expectedReturnRate, years, inflationRate)

	fmt.Println(futureValue);
	fmt.Println(futureRealValue);
}
