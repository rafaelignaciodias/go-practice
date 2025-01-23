package accounts

type CheckingAccount struct {
	Agency       int
	Account      int
	CustomerName string
	Balance      float64
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	validDeposit := depositValue > 0

	if validDeposit {
		c.Balance += depositValue
		return "Deposit successful", c.Balance
	} else {
		return "Deposit not allowed", c.Balance
	}
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {
	validWithdraw := withdrawValue > 0 && withdrawValue <= c.Balance

	if validWithdraw {
		c.Balance -= withdrawValue
		return "Withdraw successful"
	} else {
		return "Withdraw not allowed"
	}
}

func (c *CheckingAccount) Transfer(transferValue float64, destinationAccount *CheckingAccount) bool {
	validTransfer := transferValue > 0 && transferValue <= c.Balance

	if validTransfer {
		c.Balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}
