# 🚀 Release v0.1.0 — numconvert

This is the **first release** of the `numconvert` tool under **24labs-tools**. 🎉  

---

## ✨ Features
- **Automatic base detection**:
  - `0B` prefix → binary  
  - `0X` prefix → hexadecimal  
  - leading `0` → octal  
  - otherwise → decimal  

- **Conversions**:
  - `ToDecimal` (string → int64)  
  - `ToHex` (int64 → hex string)  
  - `ToBinary` (int64 → binary string, padded to 8 bits)  
  - `ToOctal` (int64 → octal string)  

- **BCD utilities**:
  - `EncodeBCD` (string of digits → BCD bytes)  
  - `DecodeBCD` (BCD bytes → string of digits)  

- **CLI tool**:
  - Run interactively with `go run ./cmd/numconvert`  
  - Example:
    ```
    Enter a number (prefix with 0B for binary, 0X for hex): 0B1010
    Decimal: 10
    Binary: 0B00001010
    Hex: 0X0A
    Octal: 012
    ```

---

## 🧪 Testing
- Unit tests cover conversion and BCD logic (~79% coverage).  
- Run:
  ```bash
  go test ./pkg/numconvert -cover
  ```

---

## 📦 Installation
```bash
go get github.com/24labs/24labs-tools/pkg/numconvert
```

---

## 📜 Notes
- This release sets the **baseline quality bar** for all future 24labs-tools.  
- Functions include full GoDoc comments and examples.  
