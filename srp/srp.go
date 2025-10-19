package main

import "fmt"

// ReportGenerator 只負責產生報表內容
type ReportGenerator struct{}

func (r *ReportGenerator) Generate() string {
    return "REPORT: data..."
}

// FileSaver 只負責儲存
type FileSaver struct{}

func (f *FileSaver) Save(filename, content string) error {
    fmt.Printf("Saving %s -> %s\n", filename, content)
    return nil
}

func main() {
    gen := &ReportGenerator{}
    saver := &FileSaver{}
    content := gen.Generate()
    _ = saver.Save("report.txt", content)
}
