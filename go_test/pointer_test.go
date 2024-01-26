package main

import "testing"

type Wallet struct{}

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {}

func (w Wallet) Balance() int {
	return 0
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
