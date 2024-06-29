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

	// inputs
	fmt.Print("Investment amount: ");
	fmt.Scan(&investmentAmount);

	fmt.Print("Years: ");
	fmt.Scan(&years);

	fmt.Print("expectedReturnRate: ");
	fmt.Scan(&expectedReturnRate);

	// calculating values
	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years);
	futureRealValue := calculateFutureValueWithInflation(investmentAmount, expectedReturnRate, years, inflationRate)

	// formatting outputs
	formattedFutureValue := fmt.Sprintf("Future value: %.2f", futureValue);
	formattedFutureRealValue := fmt.Sprintf("Future value (adjusted for inflation): %.2f", futureRealValue);

	// outputs
	fmt.Println(formattedFutureValue);
	fmt.Println(formattedFutureRealValue);
}
