# SOLID Principles in Go (Golang Examples) | Go 語言 SOLID 原則示範

This repository demonstrates **SOLID design principles** in idiomatic Go with runnable examples.  
這個專案展示了 **SOLID 物件導向設計五大原則** 的 Go 實作，每個原則都有獨立範例可執行與測試。

---

## Principles 原則

| Principle | 中文 | Description | Example |
|-----------|------|-------------|---------|
| SRP | 單一職責 | A class/module should have only one responsibility. | `srp/` |
| OCP | 開放封閉 | Open for extension, closed for modification. | `ocp/` |
| LSP | 里氏替換 | Subtypes must be substitutable for their base types. | `lsp/` |
| ISP | 介面隔離 | Clients should not be forced to depend on unused methods. | `isp/` |
| DIP | 依賴反轉 | Depend upon abstractions, not concretions. | `dip/` |

---

## Quick Start 快速開始

```bash
git clone https://github.com/yangmark/solid-example-go.git
cd solid-example-go
go test ./...
```

---
