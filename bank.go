package main

import "sync"

type Bank interface {
	deposit(amount int) bool
	withdraw(amount int) bool
}

// struct is better than a global var for thread safety and concurrecy
type Accounts struct {
	//include mutex field
	mutex sync.RWMutex
	// user_id -> balance
	user_accounts map[int]int
}

//return a reference to pointer of the new bank struct
func NewBank() *Accounts {
	return &Accounts {
		user_accounts: make(map[int]int),
	}
}

func (a *Accounts) open_account(userID int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.user_accounts[userID] = 0
}


func (a *Accounts) deposit(userID int, amount int) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	_, exists := a.user_accounts[userID]
	if exists {
		//locks access to this account while someone is depositing
		a.user_accounts[userID] += amount
		//unlocks account when the function returns
	} else {
		return false
	}
	return true
}

func (a *Accounts) withdraw(userID int, amount int) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	_, exists := a.user_accounts[userID]
	if exists && a.user_accounts[userID] >= amount {
		a.user_accounts[userID] -= amount
	} else {
		return false
	}

	return true
}

