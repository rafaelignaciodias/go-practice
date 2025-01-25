package main

import (
	"fmt"
	"intermediate/accounts"
)

func PayBill(account verifyAccount, billAmount float64) {
	account.Withdraw(billAmount)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {

	checkingAccount := accounts.Checking{}

	checkingAccount.Deposit(100)
	PayBill(&checkingAccount, 51)

	fmt.Println(checkingAccount)

	savingAccount := accounts.Saving{}

	savingAccount.Deposit(1000)
	PayBill(&savingAccount, 749)
	savingAccount.Withdraw(99)

	fmt.Println(savingAccount)
}
