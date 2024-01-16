package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay()
}

type Payment struct{}

// This principle states that a struct should be open for extension but closed for modification,
// meaning that the behavior of a struct can be extended without changing its code.
// This helps to keep the code flexible and adaptable to changing requirements.

type PayPal struct {
	amount float64
}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard", cc.amount)
}

func (pp PayPal) Pay() {
	fmt.Printf("Paid %.2f using PayPal", pp.amount)
}

func main() {
	p := Payment{}
	cc := CreditCard{12.23}
	p.Process(cc)

	//PayPal
	pp := PayPal{22.33}
	p.Process(pp)
}
