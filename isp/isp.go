package main

import "fmt"

// 分拆介面
type Printer interface {
    Print()
}

type Scanner interface {
    Scan()
}

type MultiFunctionDevice struct{}

func (m MultiFunctionDevice) Print() { fmt.Println("printing...") }
func (m MultiFunctionDevice) Scan()  { fmt.Println("scanning...") }

func main() {
    var p Printer = MultiFunctionDevice{}
    p.Print()

    var s Scanner = MultiFunctionDevice{}
    s.Scan()
}
