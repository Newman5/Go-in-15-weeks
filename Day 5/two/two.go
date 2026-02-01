// Create a struct Account with field Balance int.
// Write:
// GetBalance() (value receiver)
// Deposit(amount int) (pointer receiver)
// Verify that Deposit actually changes the balance.

package main

import "fmt"

type Account struct {
	Balance int
}

func main() {
	account := Account{Balance: 100}
	fmt.Println(account.GetBalance())
	account.Deposit(100)
	fmt.Println(account.GetBalance())
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}
