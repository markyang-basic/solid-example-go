package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    W, H float64
}
func (r Rectangle) Area() float64 { return r.W * r.H }

type Circle struct {
    R float64
}
func (c Circle) Area() float64 { return 3.14159 * c.R * c.R }

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    r := Rectangle{W: 3, H: 4}
    c := Circle{R: 2}
    PrintArea(r)
    PrintArea(c)
}
