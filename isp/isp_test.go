package main

import "testing"

func TestMFD(t *testing.T) {
    m := MultiFunctionDevice{}
    m.Print()
    m.Scan()
}
