package main

import "fmt"

type CheckingAccount struct {
	agency       int
	account      int
	customerName string
	balance      float64
}

func main() {

	regularCA := CheckingAccount{123, 456, "Regular Customer", 123.45}
	fmt.Println(regularCA)

	regularIA := CheckingAccount{agency: 321, account: 654, customerName: "Investidor Customer", balance: 543.21}
	fmt.Println(regularIA)
}
