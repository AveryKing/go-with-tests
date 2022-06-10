package main

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("Deposit wallet balance is \n %v", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
