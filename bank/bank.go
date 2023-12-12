package bank

import (
	"bank.com/files"
	"fmt"
)

const filePath = "./accounts.dat"

// CreateAccount creates a new bank account.
func CreateAccount() {
	account := files.BankAccount{}

	fmt.Print("Enter your name: ")
	_, err := fmt.Scanf("%s", &account.Name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your address: ")
	_, err = fmt.Scanf("%s", &account.Address)
	if err != nil {
		fmt.Println(err)
	}

	accountNumbers := files.GetAccountNumbers(filePath)

	if len(accountNumbers) == 0 {
		account.AccountNumber = 12345678
	} else {
		account.AccountNumber = accountNumbers[len(accountNumbers)-1] + 1
	}

	fmt.Print("Enter your account type: ")
	_, err = fmt.Scanf("%s", &account.AccountType)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your PIN: ")
	_, err = fmt.Scanf("%d", &account.Pin)
	if err != nil {
		fmt.Println(err)
	}

	files.SaveAccountToFile(filePath, account)
}

// GetAccountDetails gets the details of a bank account.
func GetAccountDetails() {
	fmt.Print("Enter your account number: ")
	var accountNumber uint64
	_, err := fmt.Scanf("%d", &accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	account := files.GetAccount(filePath, accountNumber)

	fmt.Println("Name:", account.Name)
	fmt.Println("Address:", account.Address)
	fmt.Println("Account number:", account.AccountNumber)
	fmt.Println("Account type:", account.AccountType)
	fmt.Println("Balance:", account.Balance)

}

// DepositFunds deposits funds into a bank account.
func DepositFunds() {
	fmt.Print("Enter your account number: ")
	var accountNumber uint64
	_, err := fmt.Scanf("%d", &accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your PIN: ")
	var pin uint16
	_, err = fmt.Scanf("%d", &pin)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter the amount you wish to deposit: ")
	var amount float64
	_, err = fmt.Scanf("%f", &amount)
	if err != nil {
		fmt.Println(err)
	}

	account := files.GetAccount(filePath, accountNumber)

	if account == (files.BankAccount{}) {
		fmt.Println("Account not found")
		return
	}

	if account.Pin != pin {
		fmt.Println("Invalid PIN")
		return
	}

	account.Balance += amount

	files.UpdateAccount(filePath, account)
	fmt.Println("Thank you, you're new balance is: ", account.Balance)
}

// WithdrawFunds withdraws funds from a bank account.
func WithdrawFunds() {
	fmt.Print("Enter your account number: ")
	var accountNumber uint64
	_, err := fmt.Scanf("%d", &accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your PIN: ")
	var pin uint16
	_, err = fmt.Scanf("%d", &pin)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter the amount you wish to withdraw: ")
	var amount float64
	_, err = fmt.Scanf("%f", &amount)
	if err != nil {
		fmt.Println(err)
	}

	account := files.GetAccount(filePath, accountNumber)

	if account == (files.BankAccount{}) {
		fmt.Println("Account not found")
		return
	}

	if account.Pin != pin {
		fmt.Println("Invalid PIN")
		return
	}

	if account.Balance < amount {
		fmt.Println("Insufficient funds")
		return
	}

	account.Balance -= amount

	files.UpdateAccount(filePath, account)
	fmt.Println("Thank you, you're new balance is: ", account.Balance)
}

// TransferFunds transfers funds from one bank account to another.
func TransferFunds() {
	fmt.Print("Enter your account number: ")
	var accountNumber uint64
	_, err := fmt.Scanf("%d", &accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your PIN: ")
	var pin uint16
	_, err = fmt.Scanf("%d", &pin)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter the account number you wish to transfer to: ")
	var transferAccountNumber uint64
	_, err = fmt.Scanf("%d", &transferAccountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter the amount you wish to transfer: ")
	var amount float64
	_, err = fmt.Scanf("%f", &amount)
	if err != nil {
		fmt.Println(err)
	}

	account := files.GetAccount(filePath, accountNumber)

	if account == (files.BankAccount{}) {
		fmt.Println("Account not found")
		return
	}

	if account.Pin != pin {
		fmt.Println("Invalid PIN")
		return
	}

	if account.Balance < amount {
		fmt.Println("Insufficient funds")
		return
	}

	transferAccount := files.GetAccount(filePath, transferAccountNumber)

	if transferAccount == (files.BankAccount{}) {
		fmt.Println("Account not found")
		return
	}

	account.Balance -= amount
	transferAccount.Balance += amount

	files.UpdateAccount(filePath, account)
	files.UpdateAccount(filePath, transferAccount)
	fmt.Println("Thank you, you're new balance is: ", account.Balance)
}

// RemoveAccount removes a bank account.
func RemoveAccount() {
	fmt.Print("Enter your account number: ")
	var accountNumber uint64
	_, err := fmt.Scanf("%d", &accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Enter your PIN: ")
	var pin uint16
	_, err = fmt.Scanf("%d", &pin)
	if err != nil {
		fmt.Println(err)
	}

	account := files.GetAccount(filePath, accountNumber)

	if account == (files.BankAccount{}) {
		fmt.Println("Account not found")
		return
	}

	if account.Pin != pin {
		fmt.Println("Invalid PIN")
		return
	}

	files.DeleteAccount(filePath, accountNumber)
	fmt.Println("Account deleted")
}
