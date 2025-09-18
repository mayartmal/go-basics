package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFile = "account.txt"

func main() {
	// var accountBalance float64 = readBalanceFromTheFile()
	fmt.Println("Welcome to Go Bank")
	for {
		fmt.Println("Select an action: ")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		choice := choiseSetter()
		accountBalance, err := readBalanceFromTheFile()
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			break
		}
		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			accountBalance = accountBalance + deposit(accountBalance)
			fmt.Println("Your balance is:", accountBalance)
		case 3:
			accountBalance = accountBalance - withdraw(accountBalance)
			fmt.Println("Your balance is:", accountBalance)
		case 4:
			fmt.Println("Session was terminated. Thank you.")
			return
		default:
			fmt.Println("Unknown input.")
			continue
		}
		writeBalance2TheFile(accountBalance)

		// if choice == 1 {
		// 	fmt.Println("Your balance is: ", accountBalance)
		// } else if choice == 2 {
		// 	accountBalance = accountBalance + deposit(accountBalance)
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 3 {
		// 	accountBalance = accountBalance - withdraw(accountBalance)
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 4 {
		// 	fmt.Println("Session was terminated. Thank you.")
		// 	return
		// } else {
		// 	fmt.Println("Unknown input.")
		// 	continue
		// }
	}
	fmt.Println("Googbye")

}

func writeBalance2TheFile(account float64) {
	accountText := fmt.Sprint(account)
	os.WriteFile(accountFile, []byte(accountText), 0644)
}

func readBalanceFromTheFile() (float64, error) {
	data, err := os.ReadFile(accountFile)
	if err != nil {
		return 1000, errors.New("Error during reading the file")

	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Error during parsing the file")
	}
	return balance, nil
}

func choiseSetter() int {
	var userInput int
	// fmt.Printf("Your choice is: ")
	fmt.Scan(&userInput)
	fmt.Printf("Selected operation is: %v\n", userInput)

	return userInput
}

func deposit(account float64) float64 {
	var userInput float64
	fmt.Printf("Enter deposit amount: ")
	fmt.Scan(&userInput)
	if userInput <= 0 {
		fmt.Println("Invalid amaunt")
		return 0
	}
	account += userInput
	fmt.Printf("You have %v on your  account now\n", account)
	return userInput
}

func withdraw(account float64) float64 {
	var userInput float64
	fmt.Printf("Enter withdraw amount: ")
	fmt.Scan(&userInput)
	if userInput > account || userInput <= 0 {
		fmt.Println("Invalid amaunt")
		return 0
	}
	account -= userInput
	fmt.Printf("You have %v on your  account now\n", account)
	return userInput
}
