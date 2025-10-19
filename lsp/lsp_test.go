package main

import "testing"

func TestShapes(t *testing.T) {
    shapes := []Shape{Rectangle{3, 4}, Circle{2}}
    for _, s := range shapes {
        if s.Area() <= 0 {
            t.Error("area must be positive")
        }
    }
}
