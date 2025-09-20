package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountFile = "account.txt"

func main() {
	// var accountBalance float64 = readBalanceFromTheFile()
	fmt.Println("Welcome to Go Bank")
	fmt.Println("We are located at", randomdata.City())
	interactOptions()
	for {

		choice := choiseSetter()
		accountBalance, err := fileops.GetFloatFromFile(accountFile, 1000)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			// break
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
		fileops.WriteFloat2File(accountFile, accountBalance)

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
