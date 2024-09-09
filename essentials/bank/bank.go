package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("What do you want to do?");
	fmt.Println("1. Check balance");
	fmt.Println("2. Deposit money");
	fmt.Println("3. Withdraw money");
	fmt.Println("4. Exit");

	var choice int;

	fmt.Print("Your choice: ");
	fmt.Scan(&choice);

	return choice;
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644);
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile);

	if (err != nil) {
		return 1000, errors.New("failed to find balance file");
	}

	balanceText := strings.Replace(string(data), "\n", "", -1);
	balance, err := strconv.ParseFloat(balanceText, 64);

	if (err != nil) {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil;
}

func main() {
	accountBalance, err := readBalanceFromFile();

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		return
		// panic(err);
	}

	fmt.Println("Welcome to Go Bank!");

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

		writeBalanceToFile(accountBalance)

		fmt.Print("\n")

	}
}
