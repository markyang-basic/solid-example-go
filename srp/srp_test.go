package main

import "testing"

func TestReportGenerator(t *testing.T) {
    gen := &ReportGenerator{}
    if gen.Generate() == "" {
        t.Error("expected non-empty report")
    }
}

func TestFileSaver(t *testing.T) {
    saver := &FileSaver{}
    if err := saver.Save("test.txt", "hello"); err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}
