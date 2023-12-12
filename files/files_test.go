package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var initialContent = "Account1,Here,12345678,current,0,1111\n"

// setupTestFile creates a temporary file with initialContent as its content.
func setupTestFile(t *testing.T) (filename string, cleanupFunc func()) {
	// Create a temporary test file
	tempFile, err := ioutil.TempFile("", "test-accounts.dat")
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}

	if _, err := tempFile.WriteString(initialContent); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}
	tempFile.Close()

	return tempFile.Name(), func() {
		err = os.Remove(tempFile.Name()) // Clean up
		if err != nil {
			t.Fatalf("Failed to remove temp file: %s", err)
		}
	}
}

// TestReadFile tests ReadFile with a valid file.
func TestReadFile(t *testing.T) {
	filename, cleanup := setupTestFile(t)
	defer cleanup()

	// Test ReadFile
	content, err := ReadFile(filename)
	if err != nil {
		t.Errorf("ReadFile returned an error: %s", err)
	}

	if content != initialContent {
		t.Errorf("ReadFile returned unexpected content: %s", content)
	}
}

// TestReadFileNonExistent tests ReadFile with a non-existent file.
func TestReadFileNonExistent(t *testing.T) {
	_, cleanup := setupTestFile(t)
	defer cleanup()

	_, err := ReadFile("nonexistentfile.txt")
	if err == nil {
		t.Errorf("Expected an error for non-existent file, but got none")
	}
}

// TestSaveAccountToFile tests SaveAccountToFile with a valid file.
func TestSaveAccountToFile(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	account := BankAccount{
		Name:          "Account2",
		Address:       "Here",
		AccountNumber: 12345678,
		AccountType:   "current",
		Balance:       0,
		Pin:           1111,
	}

	SaveAccountToFile(fileName, account)

	content, err := ReadFile(fileName)
	if err != nil {
		t.Errorf("ReadFile returned an error: %s", err)
	}

	expectedContent := "Account1,Here,12345678,current,0,1111\nAccount2,Here,12345678,current,0.00,1111\n"

	if content != expectedContent {
		t.Errorf("ReadFile returned unexpected content: %s", content)
	}
}

// TestSaveAccountsToFile tests SaveAccountsToFile with a valid file.
func TestSaveAccountsToFile(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	accounts := []BankAccount{
		{
			Name:          "Account2",
			Address:       "Here",
			AccountNumber: 12345678,
			AccountType:   "current",
			Balance:       0,
			Pin:           1111,
		},
		{
			Name:          "Account3",
			Address:       "Here",
			AccountNumber: 12345678,
			AccountType:   "current",
			Balance:       0,
			Pin:           1111,
		},
	}

	SaveAccountsToFile(fileName, accounts)

	content, err := ReadFile(fileName)
	if err != nil {
		t.Errorf("ReadFile returned an error: %s", err)
	}

	expectedContent := "Account2,Here,12345678,current,0.00,1111\nAccount3," +
		"Here,12345678,current,0.00,1111\n"

	if content != expectedContent {
		t.Errorf("ReadFile returned unexpected content: %s", content)
	}
}

// TestGetAccounts tests GetAccounts with a valid file.
func TestGetAccounts(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	accounts := GetAccounts(fileName)

	if accounts[0].Name != "Account1" {
		t.Errorf("Expected name to be Account1, got %s", accounts[0].Name)
	}

	if accounts[0].Address != "Here" {
		t.Errorf("Expected address to be Here, got %s", accounts[0].Address)
	}

	if accounts[0].AccountNumber != 12345678 {
		t.Errorf("Expected accountNumber to be 12345678, got %d", accounts[0].AccountNumber)
	}

	if accounts[0].AccountType != "current" {
		t.Errorf("Expected accountType to be current, got %s", accounts[0].AccountType)
	}

	if accounts[0].Balance != 0 {
		t.Errorf("Expected balance to be 0, got %f", accounts[0].Balance)
	}

	if accounts[0].Pin != 1111 {
		t.Errorf("Expected pin to be 1111, got %d", accounts[0].Pin)
	}
}

func TestGetAccount(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	account := GetAccount(fileName, 12345678)

	if account.Name != "Account1" {
		t.Errorf("Expected name to be Account1, got %s", account.Name)
	}

	if account.Address != "Here" {
		t.Errorf("Expected address to be Here, got %s", account.Address)
	}

	if account.AccountNumber != 12345678 {
		t.Errorf("Expected accountNumber to be 12345678, got %d", account.AccountNumber)
	}

	if account.AccountType != "current" {
		t.Errorf("Expected accountType to be current, got %s", account.AccountType)
	}

	if account.Balance != 0 {
		t.Errorf("Expected balance to be 0, got %f", account.Balance)
	}

	if account.Pin != 1111 {
		t.Errorf("Expected pin to be 1111, got %d", account.Pin)
	}
}

// TestGetAccountNumbers tests GetAccountNumbers with a valid file.
func TestGetAccountNumbers(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	accountNumbers := GetAccountNumbers(fileName)

	if len(accountNumbers) != 1 {
		t.Errorf("Expected 1 account number, got %d", len(accountNumbers))
	}

	if accountNumbers[0] != 12345678 {
		t.Errorf("Expected account number to be 12345678, got %d", accountNumbers[0])
	}
}

// TestUpdateAccount tests UpdateAccount with a valid file.
func TestUpdateAccount(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	account := BankAccount{
		Name:          "Account2",
		Address:       "Here",
		AccountNumber: 12345678,
		AccountType:   "current",
		Balance:       200,
		Pin:           1111,
	}

	UpdateAccount(fileName, account)

	accounts := GetAccounts(fileName)

	if accounts[0].Name != "Account2" {
		t.Errorf("Expected name to be Account2, got %s", accounts[0].Name)
	}

	if accounts[0].Address != "Here" {
		t.Errorf("Expected address to be Here, got %s", accounts[0].Address)
	}

	if accounts[0].AccountNumber != 12345678 {
		t.Errorf("Expected accountNumber to be 12345678, got %d", accounts[0].AccountNumber)
	}

	if accounts[0].AccountType != "current" {
		t.Errorf("Expected accountType to be current, got %s", accounts[0].AccountType)
	}

	if accounts[0].Balance != 200 {
		t.Errorf("Expected balance to be 200, got %f", accounts[0].Balance)
	}

	if accounts[0].Pin != 1111 {
		t.Errorf("Expected pin to be 1111, got %d", accounts[0].Pin)
	}
}

// TestDeleteAccount tests DeleteAccount with a valid file.
func TestDeleteAccount(t *testing.T) {
	fileName, cleanup := setupTestFile(t)
	defer cleanup()

	DeleteAccount(fileName, 12345678)

	accounts := GetAccounts(fileName)

	fmt.Println("accounts", accounts)

	if len(accounts) != 0 {
		t.Errorf("Expected 0 accounts, got %d", len(accounts))
	}
}
