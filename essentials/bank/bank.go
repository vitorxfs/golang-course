package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"

	"example.com/bank/fileops"
)

const accountBalanceFile = "data/balance.txt";

func deposit(accountBalance *float64) {
	var amount float64;
	fmt.Print("How much money do you want to deposit? ");
	fmt.Scan(&amount);
	*accountBalance += amount;
	fmt.Println("You now have ", *accountBalance);
}

func withdraw(accountBalance *float64) {
	var amount float64;
	fmt.Print("How much money do you want to withdraw? ");
	fmt.Scan(&amount);
	*accountBalance -= amount;
	fmt.Println("You now have ", *accountBalance);
}

func getOption() int {
	presentOptions();

	var choice int;

	fmt.Print("Your choice: ");
	fmt.Scan(&choice);

	return choice;
}

func main() {
	accountBalance, err := fileops.ReadFloatFromFile(accountBalanceFile);

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		return
		// panic(err);
	}

	fmt.Println("Welcome to Go Bank!");
	fmt.Println("Reach us 24/7: ", randomdata.PhoneNumber())

	for {
		choice := getOption();

		switch choice {
			case 1:
				fmt.Println("Your balance is ", accountBalance);
			case 2:
				deposit(&accountBalance);
			case 3:
				withdraw(&accountBalance);
			default:
				return;
		}

		fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		fmt.Print("\n")
	}
}
