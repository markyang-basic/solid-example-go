package main

import "testing"

type MockNotifier struct {
    called bool
}

func (m *MockNotifier) Notify(message string) error {
    m.called = true
    return nil
}

func TestOrderService(t *testing.T) {
    mock := &MockNotifier{}
    svc := NewOrderService(mock)
    svc.CreateOrder(1)
    if !mock.called {
        t.Error("expected notifier to be called")
    }
}
