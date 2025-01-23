package main

import "fmt"

type CheckingAccount struct {
	agency       int
	account      int
	customerName string
	balance      float64
}

func (c *CheckingAccount) withdraw(withdrawValue float64) string {
	validWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if validWithdraw {
		c.balance -= withdrawValue
		return "Withdraw successful"
	} else {
		return "Withdraw not allowed"
	}
}

func main() {
	regularCA := CheckingAccount{123, 456, "Regular Customer", 123.}
	fmt.Println(regularCA)

	fmt.Println(regularCA.withdraw(1000.0))
	fmt.Println(regularCA)
}
