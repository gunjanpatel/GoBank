package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposi must be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("negative amount is not allowed to withdraw")
	}

	if amount > a.Balance {
		return errors.New("not allowed to withdraw more than you have")
	}

	a.Balance -= amount
	return nil
}

func (fromAccount *Account) Transfer(amount float64, toAccount *Account) error {

	if err := fromAccount.Withdraw(amount); err != nil {
		return err
	}

	if err := toAccount.Deposit(amount); err != nil {
		return err
	}

	return nil
}

// Default Statement method
func (a Account) Statement() string {
	return fmt.Sprintf("%v - %v - %.2f DKK", a.Number, a.Name, a.Balance)
}
