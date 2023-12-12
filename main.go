package main

import (
	"bank.com/bank"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the bank!")
	fmt.Println("How can we help you today?")

	fmt.Println("Please choose an option:")
	fmt.Println("1. Create Account")
	fmt.Println("2. Get Account Details")
	fmt.Println("3. Deposit Funds")
	fmt.Println("4. Withdraw Funds")
	fmt.Println("5. Transfer Funds")
	fmt.Println("6. Delete Account")
	fmt.Println("7. Exit")

	i := 0

	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You entered", i)

	if i == 1 {
		bank.CreateAccount()
	}
	if i == 2 {
		bank.GetAccountDetails()
	}
	if i == 3 {
		bank.DepositFunds()
	}
	if i == 4 {
		bank.WithdrawFunds()
	}
	if i == 5 {
		bank.TransferFunds()
	}
	if i == 6 {
		bank.RemoveAccount()
	}
}
