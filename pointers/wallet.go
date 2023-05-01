package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(cash Bitcoin) {
	fmt.Print(w.balance)
	w.balance += cash
}

var ErrorInsuffFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(cash Bitcoin) error {
	if w.balance < cash {
		return ErrorInsuffFunds
	}
	w.balance -= cash
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
