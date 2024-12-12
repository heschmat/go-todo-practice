package account

// **Account** defines behaviors for any Type of account
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	CheckBalance()
	PrintTransactions()
}
