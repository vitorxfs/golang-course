package main

import (
	"fmt"
	"os"
)

const profitFile = "data/profit.txt";

func calculateEBT(revenue float64, expenses float64) float64 {
	return revenue - expenses;
}

func calculateProfit(ebt float64, taxRate float64) float64 {
	return ebt * (1 - taxRate/100);
}

func calculateRatio(ebt float64, profit float64) float64 {
	return ebt / profit;
}

func valid(value float64) bool {
	return value > 0
}

func writeResultsTofile(ebt, profit, ratio float64) error {
	formattedText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio);
	return os.WriteFile(profitFile, []byte(formattedText), 0644);
}

func main() {
	var revenue, expenses, taxRate float64;

	fmt.Print("Revenue: ");
	fmt.Scan(&revenue);
	if (!valid(revenue)) {
		fmt.Println("revenue input is invalid");
		return;
	}

	fmt.Print("Expenses: ");
	fmt.Scan(&expenses);
	if (!valid(expenses)) {
		fmt.Println("expenses input is invalid");
		return;
	}

	fmt.Print("Tax Rate (%): ");
	fmt.Scan(&taxRate);
	if (!valid(taxRate)) {
		fmt.Println("tax rate input is invalid");
		return;
	}

	ebt := calculateEBT(revenue, expenses);
	profit := calculateProfit(ebt, taxRate);
	ratio := calculateRatio(ebt, profit);

	fmt.Println("EBT: ", ebt);
	fmt.Println("Profit: ", profit);
	fmt.Println("Ratio: ", ratio);

	error := writeResultsTofile(ebt, profit, ratio);
	if (error != nil) {
		panic(error);
	}
}
