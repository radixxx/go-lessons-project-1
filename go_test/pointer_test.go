package main

import (
	"testing"
)

type Wallet struct{}

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(100)

	got := wallet.Balance()
	want := 100

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
