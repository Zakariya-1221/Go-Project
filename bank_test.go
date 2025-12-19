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

func TestTransfer(t *testing.T) {
	bank := NewBank()
	user1 := 10
	user2 := 20

	// Setup: Open accounts for both users
	bank.open_account(user1)
	bank.open_account(user2)
	bank.deposit(user1, 500) // User 1 starts with $500
	bank.deposit(user2, 100) // User 2 starts with $100

	t.Run("Successful transfer between two accounts", func(t *testing.T) {
		transferAmount := 200
		
		// Capture balances before the move
		u1Initial := bank.user_accounts[user1]
		u2Initial := bank.user_accounts[user2]

		got := bank.transfer(user1, user2, transferAmount)

		if got != true {
			t.Fatalf("Transfer failed: expected true, got %t", got)
		}

		// Verify User 1 was debited
		if bank.user_accounts[user1] != u1Initial-transferAmount {
			t.Errorf("Sender balance incorrect: got %d, want %d", bank.user_accounts[user1], u1Initial-transferAmount)
		}

		// Verify User 2 was credited
		if bank.user_accounts[user2] != u2Initial+transferAmount {
			t.Errorf("Receiver balance incorrect: got %d, want %d", bank.user_accounts[user2], u2Initial+transferAmount)
		}
	})

	t.Run("Transfer fails due to insufficient funds", func(t *testing.T) {
		// User 1 currently has $300 (from previous subtest)
		u1Before := bank.user_accounts[user1]
		u2Before := bank.user_accounts[user2]

		got := bank.transfer(user1, user2, 1000) // Try to send $1000

		if got != false {
			t.Errorf("Expected false for insufficient funds, but got %t", got)
		}

		// Verify NO balances changed
		if bank.user_accounts[user1] != u1Before || bank.user_accounts[user2] != u2Before {
			t.Errorf("Balances changed during failed transfer! U1: %d, U2: %d", bank.user_accounts[user1], bank.user_accounts[user2])
		}
	})

	t.Run("Transfer fails if recipient does not exist", func(t *testing.T) {
		nonExistentUser := 999
		got := bank.transfer(user1, nonExistentUser, 50)

		if got != false {
			t.Errorf("Expected false when transferring to non-existent user, but got %t", got)
		}
	})
}