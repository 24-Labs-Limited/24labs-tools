# 🛠 24labs-tools

A collection of reusable **Go utilities** designed for **fintech, payments, and systems programming**.  
Each tool is structured as a library under `pkg/` with a corresponding CLI under `cmd/`.  

This repo is actively developed and will expand to cover **ISO 8583 message processing, BCD utilities, Transact TPH helpers**, and more.  

---

## 📂 Structure

```
24labs-tools/
│
├── pkg/            # Reusable libraries
│   ├── numconvert/ # Numeric base conversions + BCD
│   ├── bcdutil/    # (planned) Dedicated BCD utilities
│   ├── tphutil/    # (planned) Transact TPH helpers
│   └── ...
│
├── cmd/            # CLI entrypoints
│   ├── numconvert/
│   ├── bcdutil/
│   └── tphutil/
│
├── ROADMAP.md      # High-level roadmap of tools
├── CHANGELOG.md    # Top-level changelog
└── README.md       # This file
```

---

## ✨ Current Tools

### 🔢 numconvert
Numeric base conversions (binary, decimal, octal, hexadecimal) + **BCD utilities**.  
- Detect base automatically (prefix-aware).  
- Convert between bases.  
- Encode/decode BCD.  
- Includes CLI:  
  ```bash
  go run ./cmd/numconvert
  ```
  Example:
  ```
  Enter a number (prefix with 0B for binary, 0X for hex): 0B1010
  Decimal: 10
  Binary: 0B00001010
  Hex: 0X0A
  Octal: 012
  ```

📖 See [pkg/numconvert/README.md](./pkg/numconvert/README.md) for details.  

---

## 📍 Roadmap
- **bcdutil** → Standalone BCD encoder/decoder with advanced options.  
- **tphutil** → Tools for parsing and building Transact TPH structures.  
- **iso8583util** → Message parsing/validation for ISO 8583.  
- Additional fintech/system utilities to be added iteratively.  

See [ROADMAP.md](./ROADMAP.md) for the full plan.  

---

## 🧪 Testing
Run all tests:  
```bash
go test ./... -cover
```

Run specific tool tests:  
```bash
go test ./pkg/numconvert -v
```

---

## 📦 Installation
Each tool can be imported individually, e.g.:  
```bash
go get github.com/24labs/24labs-tools/pkg/numconvert
```

---

## 📜 License
MIT License © 24labs
