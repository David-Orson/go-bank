package files

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BankAccount struct {
	Name          string
	Address       string
	AccountNumber uint64
	AccountType   string
	Balance       float64
	Pin           uint16
}

// ReadFile reads a file and returns its content as a string.
func ReadFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	res := string(file)

	return res, nil
}

// SaveAccountToFile saves a BankAccount to a file.
func SaveAccountToFile(filename string, account BankAccount) {
	accountData := account.Name + "," + account.Address + "," + strconv.FormatUint(
		account.AccountNumber, 10,
	) + "," + account.AccountType + "," + strconv.FormatFloat(account.Balance, 'f', 2, 64) + "," + strconv.FormatUint(
		uint64(account.Pin), 10,
	) + "\n"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("Failed to close file:", err)
			return
		}
	}()

	_, err = file.WriteString(accountData)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}
}

// SaveAccountsToFile saves a slice of BankAccounts to a file.
func SaveAccountsToFile(filename string, accounts []BankAccount) {
	if len(accounts) == 0 {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println("Failed to remove file:", err)
			return
		}
	}

	fileString := ""

	for _, account := range accounts {
		accountData := account.Name + "," + account.Address + "," + strconv.FormatUint(
			account.AccountNumber, 10,
		) + "," + account.AccountType + "," + strconv.FormatFloat(
			account.Balance, 'f', 2, 64,
		) + "," + strconv.FormatUint(
			uint64(account.Pin), 10,
		) + "\n"

		fileString += accountData
	}

	_, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("Failed to close file:", err)
			return
		}
	}()

	_, err = file.WriteString(fileString)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}
}

// GetAccounts reads a file and returns its content as a BankAccount.
func GetAccounts(fileName string) []BankAccount {
	var accounts []BankAccount

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return accounts
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return accounts
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("Failed to close file:", err)
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		accountData := strings.Split(line, ",")
		accountNumber, err := strconv.ParseUint(accountData[2], 10, 64)
		if err != nil {
			fmt.Println("Failed to parse account number:", err)
			return accounts
		}

		balance, err := strconv.ParseFloat(accountData[4], 64)
		if err != nil {
			fmt.Println("Failed to parse balance:", err)
			return accounts
		}

		pin, err := strconv.ParseUint(accountData[5], 10, 16)
		if err != nil {
			fmt.Println("Failed to parse pin:", err)
			return accounts
		}

		account := BankAccount{
			Name:          accountData[0],
			Address:       accountData[1],
			AccountNumber: accountNumber,
			AccountType:   accountData[3],
			Balance:       balance,
			Pin:           uint16(pin),
		}

		accounts = append(accounts, account)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan file:", err)
		return accounts
	}

	return accounts
}

// GetAccount reads a file and returns its content as a BankAccount.
func GetAccount(fileName string, accountNumber uint64) BankAccount {
	accounts := GetAccounts(fileName)

	for _, account := range accounts {
		if account.AccountNumber == accountNumber {
			return account
		}
	}

	return BankAccount{}
}

// GetAccountNumbers reads a file and returns the bank account numbers in that file.
func GetAccountNumbers(fileName string) []uint64 {
	accounts := GetAccounts(fileName)
	var accountNumbers []uint64

	for _, account := range accounts {
		accountNumbers = append(accountNumbers, account.AccountNumber)
	}

	return accountNumbers
}

// UpdateAccount updates a BankAccount in a file.
func UpdateAccount(fileName string, account BankAccount) {
	accounts := GetAccounts(fileName)

	for i, acc := range accounts {
		if acc.AccountNumber == account.AccountNumber {
			accounts[i] = account
		}
	}

	SaveAccountsToFile(fileName, accounts)
}

// DeleteAccount deletes a BankAccount from a file.
func DeleteAccount(fileName string, accountNumber uint64) {
	accounts := GetAccounts(fileName)

	for i, account := range accounts {
		if account.AccountNumber == accountNumber {
			accounts = append(accounts[:i], accounts[i+1:]...)
		}
	}
	SaveAccountsToFile(fileName, accounts)
}
