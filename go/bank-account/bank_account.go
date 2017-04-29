// API:
//
// Open(initialDeposit int64) *Account
// (Account) Close() (payout int64, ok bool)
// (Account) Balance() (balance int64, ok bool)
// (Account) Deposit(amount int64) (newBalance int64, ok bool)
//
// If Open is given a negative initial deposit, it must return nil.
// Deposit must handle a negative amount as a withdrawal.
// If any Account method is called on an closed account, it must not modify
// the account and must return ok = false.

package account

import "sync"

const testVersion = 1

type Account struct {
	sync.Mutex
	active  bool
	balance int64
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		active:  true,
		balance: initialDeposit,
	}
}

func (account *Account) Close() (int64, bool) {
	account.Lock()
	defer account.Unlock()
	if !account.active {
		return 0, false
	}
	account.active = false
	return account.balance, true
}

func (account *Account) Balance() (int64, bool) {
	account.Lock()
	defer account.Unlock()
	if !account.active {
		return 0, false
	}
	return account.balance, true
}

func (account *Account) Deposit(amount int64) (int64, bool) {
	account.Lock()
	defer account.Unlock()
	if !account.active {
		return 0, false
	}
	if account.balance+amount < 0 {
		return account.balance, false
	}
	account.balance += amount
	return account.balance, true
}
