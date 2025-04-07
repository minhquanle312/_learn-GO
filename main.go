package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.0

	for {
		fmt.Println("Welcome to Go Bank")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is: $", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated, new amount: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount < 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You cannot withdraw more than your balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated, new amount: ", accountBalance)
		} else {
			fmt.Println("Exiting the program.")
			break
		}
	}

	fmt.Println("Thank you for using Go Bank!")
}
