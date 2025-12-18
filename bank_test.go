package main

import (
	"testing"
)

func TestDeposit(t *testing.T){
	bank_accountA := NewBank()
	userID := 1
	amount := 100

	t.Run("Open a new account for a user", func(t *testing.T) {
		bank_accountA.open_account(userID)
		balance, exists := bank_accountA.user_accounts[userID]
		if !exists {
			t.Fatalf("Expected account %d to exist, but does not", userID)
		}

		if balance != 0 {
			t.Fatalf("Invalid balance: Got balance of %d, expected 0", balance)
		}
	})

	t.Run("Depositing to a new user account", func(t *testing.T) {
		initialBalance := bank_accountA.user_accounts[userID]
		got := bank_accountA.deposit(userID, amount)
		want := true
		if got != want {
			t.Fatalf("Got %t, want %t", got, want)
		}

		newBalance := bank_accountA.user_accounts[userID]
		expectedBalance := initialBalance + amount
		if newBalance != expectedBalance {
			t.Errorf("Invalid account balance after deposit, got: %d, expected %d", newBalance, expectedBalance)
		}
	})

}

func TestWithdraw(t *testing.T) {
	bank_accountB := NewBank()
	userID := 2
	amount := 75
	t.Run("Withdraw from a user account with sufficient amount", func(t *testing.T) {
		//don't need to check for existence since previous run tested for that
		bank_accountB.open_account(userID)
		bank_accountB.deposit(userID, 100) //start with 100
		initialBalance := bank_accountB.user_accounts[userID]
		got := bank_accountB.withdraw(userID, amount)
		newBalance := bank_accountB.user_accounts[userID]
		expectedBalance := initialBalance - amount
		if got != true {
			t.Fatalf("Got %t and amount %d, want %t and expected %d", got, newBalance, true, expectedBalance)
		}

	})

	t.Run("Withdraw from a user account with insufficient amount", func(t *testing.T) {
		//don't need to check for existence since previous run tested for that
		initialBalance := bank_accountB.user_accounts[userID]
		got := bank_accountB.withdraw(userID, 1000000)
		newBalance := bank_accountB.user_accounts[userID]
		if got != false {
			t.Fatalf("Invalid withdrawal, got: %t, account only had %d, expected: %t, and %d", got, initialBalance, false, newBalance)
		}

	})
}