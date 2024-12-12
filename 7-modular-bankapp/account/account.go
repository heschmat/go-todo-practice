package account

import (
	"errors"
	"fmt"
)

// `Transaction` represents a single transaction
type Transaction struct {
	Type   string  // "Deposit" or "Withdrawal"
	Amount float64
	Status string  // "Success" or "Famiald"
}


// BankAccount represents a simple bank account
type BankAccount struct {
	AccountNumber string
	Balance		  float64
	Transactions   []Transaction
}

// handleTransaction logs the transaction and sets the status.
func (b *BankAccount) handleTransaction(t *Transaction, success bool, errMsg string) error {
	if success {
		t.Status = "OK"
	} else {
		t.Status = "XX"
	}

	b.Transactions = append(b.Transactions, *t)

	if !success {
		return errors.New(errMsg)
	}
	return nil
}


// Deposit adds money to the account and logs the transaction
func (b *BankAccount) Deposit(amount float64) error {
	transaction := &Transaction{Type: "Deposit", Amount: amount}
	if !isAmountPositive(amount) {
		return b.handleTransaction(transaction, false, "deposit amount must be positive!")
	}

	b.Balance += amount
	return b.handleTransaction(transaction, true, "")
}


// Withdraw subtracts money from the account and logs the transaction
func (b *BankAccount) Withdraw(amount float64) error {
	transaction := &Transaction{Type: "Withdrawal", Amount: amount}

	if !isAmountPositive(amount) {
		return b.handleTransaction(transaction, false, "Withdrawal amount must be positive.")
	}

	if amount > b.Balance {
		return b.handleTransaction(transaction, false, "insufficient funds")
	}

	b.Balance -= amount
	return b.handleTransaction(transaction, true, "")
}


// CheckBalance displays the current balance
func (b BankAccount) CheckBalance() {
	fmt.Println("# ------------------------------------ #")
	fmt.Printf("# Account: %s ----------------- #\n", b.AccountNumber)
	fmt.Printf(" *** balance: $%.2f ***\n", b.Balance)
	fmt.Println("# ==================================== #")
}

// PrintTransactions prints the transaction history
func (b BankAccount) PrintTransactions() {
	fmt.Println("# ------ Transaction History: -------- #")
	fmt.Printf(">>> AccountNr: %s\n", b.AccountNumber)
	for _, t := range b.Transactions {
		fmt.Printf("%-10s: $%.2f - %s\n", t.Type, t.Amount, t.Status)
	}
	fmt.Println("# ==================================== #")
}
