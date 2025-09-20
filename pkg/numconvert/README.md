# 🔢 numconvert

`numconvert` is a Go library and CLI for working with **numeric base conversions** and **BCD (Binary-Coded Decimal)**.

---

## ✨ Features
- Detect numeric base automatically (prefix-aware, with safe fallback).
- Convert numbers between:
  - Decimal
  - Binary
  - Hexadecimal
  - Octal
- Encode and decode **BCD** values.
- CLI tool for quick conversions.

---

## 🧮 Base Detection Rules

`DetectBase` follows these rules:

1. **Prefix-based detection**
   - `0B` → Binary  
   - `0X` → Hexadecimal  

2. **Octal detection (Safe mode)**  
   - If a number starts with `0` and has more than one digit:  
     - It is treated as **octal** only if all characters are between `0–7`.  
     - Example: `0123` → octal  
     - Example: `0777` → octal  
     - Example: `0Z123` → **decimal** (invalid octal digit → fallback)  

3. **Hexadecimal fallback**  
   - If no prefix is present, but the string contains any `A–F` characters, it is treated as **hexadecimal**.  
   - Example: `1A2F` → hex  

4. **Default**  
   - If none of the above apply, the number is treated as **decimal**.  
   - Example: `123` → decimal  

---

## 📦 Installation

Library usage:
```bash
go get github.com/24labs/24labs-tools/pkg/numconvert
```

CLI usage (from repo):
```bash
go run ./cmd/numconvert
```

---

## 🚀 Examples

### CLI
```bash
Enter a number (prefix with 0B for binary, 0X for hex): 0B1010
Decimal: 10
Binary: 0B00001010
Hex: 0X0A
Octal: 012
```

### Library
```go
import "github.com/24labs/24labs-tools/pkg/numconvert"

val, _ := numconvert.ToDecimal("0X1A", 16)
// val == 26

hex := numconvert.ToHex(26)
// hex == "0X1A"

oct := numconvert.ToOctal(83)
// oct == "0123"
```

---

## 🧪 Testing

Run unit tests:
```bash
go test ./pkg/numconvert -v -cover
```

---

## 📜 License
MIT License © 24labs
