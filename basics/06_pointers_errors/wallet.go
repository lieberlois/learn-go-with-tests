package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is the currency
type Bitcoin int

// Wallet holds the state of a wallet
type Wallet struct {
    balance Bitcoin
}

// Stringer is an interface for String() purposes which is available by default
type Stringer interface {
	String() string
}

// ErrInsufficientFunds holds the error message as a single source of truth
var ErrInsufficientFunds = errors.New("insufficient funds")

// Deposit puts an amount of money onto a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

// Withdraw removes an amount of money from a wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance - amount < 0{
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance returns a wallets balance
func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

// String implementation for Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}



