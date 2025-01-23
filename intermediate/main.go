package main

import (
	"fmt"
	"intermediate/accounts"
)

func main() {
	regularCA := accounts.CheckingAccount{123, 456, "Regular Customer", 1000.}
	fmt.Println("Creation successful")
	fmt.Println(regularCA)

	status, _ := regularCA.Deposit(1000.0)
	fmt.Println()
	fmt.Println(status)
	fmt.Println(regularCA)

	fmt.Println()
	fmt.Println(regularCA.Withdraw(1000.0))
	fmt.Println(regularCA)

	fmt.Println()
	investimentCA := accounts.CheckingAccount{321, 654, "Investiment Customer", 1000.}
	fmt.Println("Creation successful")
	fmt.Println(investimentCA)

	fmt.Println()
	regularCA.Transfer(500., &investimentCA)
	fmt.Println("Transfer successful")
	fmt.Println(regularCA)
	fmt.Println(investimentCA)
}
