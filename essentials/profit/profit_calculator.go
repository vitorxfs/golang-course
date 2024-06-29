package main

import "fmt"

func calculateEBT(revenue float64, expenses float64) float64 {
	return revenue - expenses;
}

func calculateProfit(ebt float64, taxRate float64) float64 {
	return ebt * (1 - taxRate/100);
}

func calculateRatio(ebt float64, profit float64) float64 {
	return ebt / profit;
}

func main() {
	var revenue, expenses, taxRate float64;

	fmt.Print("Revenue: ");
	fmt.Scan(&revenue);

	fmt.Print("Expenses: ");
	fmt.Scan(&expenses);

	fmt.Print("Tax Rate (%): ");
	fmt.Scan(&taxRate);

	ebt := calculateEBT(revenue, expenses);
	profit := calculateProfit(ebt, taxRate);
	ratio := calculateRatio(ebt, profit);

	fmt.Println("EBT: ", ebt);
	fmt.Println("Profit: ", profit);
	fmt.Println("Ratio: ", ratio);
}
