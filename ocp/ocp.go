package main

import "fmt"

// Payment 抽象介面
type Payment interface {
    Pay(amount int) error
}

type CreditCard struct{}
func (c *CreditCard) Pay(amount int) error {
    fmt.Println("CreditCard pay", amount)
    return nil
}

type PayPal struct{}
func (p *PayPal) Pay(amount int) error {
    fmt.Println("PayPal pay", amount)
    return nil
}

type Checkout struct {
    payer Payment
}

func (c *Checkout) Process(amount int) error {
    return c.payer.Pay(amount)
}

func main() {
    cc := &CreditCard{}
    co := &Checkout{payer: cc}
    co.Process(100)

    pp := &PayPal{}
    co.payer = pp
    co.Process(200)
}
