# 24labs-tools 🧰

A growing collection of small, modular Go utilities ("Lego bricks") built and maintained by [24labs](https://github.com/24labs).  

Each tool is designed to be:
- **Reusable** — clean packages under `pkg/`
- **Accessible** — thin CLI wrappers under `cmd/`
- **Educational** — built step-by-step to deepen Go expertise

---

## 📦 Structure

```
24labs-tools/
├── cmd/           # CLI entrypoints
│   └── numconvert # Example: number base converter CLI
│       └── main.go
│
└── pkg/           # Reusable libraries
    └── numconvert # Example: number base converter package
        └── convert.go
```

- **`pkg/`** — reusable Go packages  
- **`cmd/`** — small CLI tools that call into `pkg/`

---

## 🧱 Available Tools

### 🔢 numconvert
A number base conversion utility.

- Input a number in **binary (`0B...`)**, **hex (`0X...`)**, or **decimal**
- Outputs the number in **decimal, binary, and hex**

#### Example

```bash
go run cmd/numconvert/main.go
```

```
Enter a number (prefix with 0B for binary, 0X for hex): 42
Decimal: 42
Binary: 0B00101010
Hex: 0X2A
```

---

## 🚀 Getting Started

Clone the repo:

```bash
git clone git@github.com:24labs/24labs-tools.git
cd 24labs-tools
```

Run any tool:

```bash
go run cmd/numconvert/main.go
```

Import a library in your Go project:

```go
import "github.com/24labs/24labs-tools/pkg/numconvert"
```

---

## 🌱 Roadmap

- Add more small, composable Go utilities
- Unit tests for each package
- CI with GitHub Actions

---

## 📜 License

MIT © [24labs](https://github.com/24labs)
