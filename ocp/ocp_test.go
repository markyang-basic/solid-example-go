package main

import "testing"

func TestCheckoutWithCreditCard(t *testing.T) {
    co := &Checkout{payer: &CreditCard{}}
    if err := co.Process(100); err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}
