package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Gunjan",
			Address: "Test 123",
			Phone:   "12345678",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("Can't create an Account objecrt")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Gunjan",
			Address: "Test 123",
			Phone:   "12345678",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("failed to add deposit in account")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Gunjan",
			Address: "Test 123",
			Phone:   "12345678",
		},
		Number:  1001,
		Balance: 0,
	}

	// We expect error here - that means our test should fail when err is nil
	if err := account.Deposit(-10); err == nil {
		t.Error("only positive number is allowed to deposite")
	}

	if err := account.Deposit(0); err == nil {
		t.Error("only positive number is allowed to deposite")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Gunjan",
			Address: "Test 123",
			Phone:   "12345678",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(5)

	if account.Balance != 5 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestWithdrawInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Gunjan",
			Address: "Test 123",
			Phone:   "12345678",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if err := account.Withdraw(-10); err == nil {
		t.Error("negative withdraw should not be allowed")
	}

	if err := account.Withdraw(0); err == nil {
		t.Error("zero amount should not be allowed to withdraw")
	}

	if err := account.Withdraw(12); err == nil {
		t.Error("should not allow to withdraw more than account has")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Gunjan",
			Address: "Test 123",
			Phone:   "12345678",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)

	if statement := account.Statement(); statement != "1001 - Gunjan - 100.00 DKK" {
		t.Error("statement doesn't have the proper format")
	}
}
