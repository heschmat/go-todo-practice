package main

import (
	"bankapp/account"
	"fmt"
	"log"
)

func main() {
	// Create a logger for the application
	log.SetPrefix("[BankApp] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Create a new bank account
	acc1 := &account.BankAccount {
		AccountNumber: "811202069",
		Balance: 500.0,
	}

	log.Println("Starting transactions for Account 1")

	acc1.CheckBalance() // Check initial balance

	// Attempt to withdraw more than the balance
	if err := acc1.Withdraw(700); err != nil {
		log.Println("Withdrawal failed:", err)
	}

	// Deposit money
	if err := acc1.Deposit(200); err != nil {
		log.Println(("Deposit failed"), err)
	} else {
		log.Println("Deposit successful")
	}

	// Withdraw a valid amount
	if err := acc1.Withdraw(700); err != nil {
		log.Println("Withdrawal failed:", err)
	}

	// Final balance check
	acc1.CheckBalance()

	// Print transaction history
	acc1.PrintTransactions()

	// ------------- create another bank account ------------- //
	// acc2 := account.BankAccount{
	// 	AccountNumber: "871926734",
	// 	Balance: 300.0,
	// }

	// log.Println("Starting transactions for Account 2")

	// acc2.CheckBalance()
	// if err := acc2.Deposit(100); err != nil {
	// 	log.Println("Deposit failed:", err)
	// }

	// acc2.CheckBalance()

	// // Print transaction history for the 2nd account:
	// acc2.PrintTransactions()

	fmt.Println("BankApp finished executing.")
}
