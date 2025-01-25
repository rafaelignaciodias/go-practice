package accounts

import "intermediate/customer"

type Checking struct {
	Customer        customer.Titular
	Agency, Account int
	balance         float64
}

func (c *Checking) Deposit(depositValue float64) (string, float64) {
	validDeposit := depositValue > 0

	if validDeposit {
		c.balance += depositValue
		return "Deposit successful", c.balance
	} else {
		return "Deposit not allowed", c.balance
	}
}

func (c *Checking) Withdraw(withdrawValue float64) string {
	validWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if validWithdraw {
		c.balance -= withdrawValue
		return "Withdraw successful"
	} else {
		return "Withdraw not allowed"
	}
}

func (c *Checking) Transfer(transferValue float64, destinationAccount *Checking) bool {
	validTransfer := transferValue > 0 && transferValue <= c.balance

	if validTransfer {
		c.balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *Checking) GetBalance() float64 {
	return c.balance
}
