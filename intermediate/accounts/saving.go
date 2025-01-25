package accounts

import "intermediate/customer"

type Saving struct {
	Customer        customer.Titular
	Agency, Account int
	Operation       int
	balance         float64
}

func (c *Saving) Deposit(depositValue float64) (string, float64) {
	validDeposit := depositValue > 0

	if validDeposit {
		c.balance += depositValue
		return "Deposit successful", c.balance
	} else {
		return "Deposit not allowed", c.balance
	}
}

func (c *Saving) Withdraw(withdrawValue float64) string {
	validWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if validWithdraw {
		c.balance -= withdrawValue
		return "Withdraw successful"
	} else {
		return "Withdraw not allowed"
	}
}
