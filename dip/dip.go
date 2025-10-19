package main

import "fmt"

// 抽象
type Notifier interface {
    Notify(message string) error
}

// 低階實作
type EmailNotifier struct{}
func (e *EmailNotifier) Notify(message string) error {
    fmt.Println("send email:", message)
    return nil
}

// 高階模組
type OrderService struct {
    notifier Notifier
}

func NewOrderService(n Notifier) *OrderService {
    return &OrderService{notifier: n}
}

func (s *OrderService) CreateOrder(id int) {
    _ = s.notifier.Notify(fmt.Sprintf("order %d created", id))
}

func main() {
    email := &EmailNotifier{}
    svc := NewOrderService(email)
    svc.CreateOrder(42)
}
