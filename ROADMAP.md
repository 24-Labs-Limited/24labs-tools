# 24labs-tools Roadmap

## Phase 1 — Foundation
- ✅ **numconvert**
  - Decimal ↔ Binary ↔ Hex ↔ Octal
  - BCD Encode (done) / Decode (next)
  - Error handling improvements
  - Unit tests + README update

---

## Phase 2 — ISO8583 Utilities
- **iso8583utils**
  - Hex bitmap → binary
  - Print active fields
  - CLI under `cmd/iso8583utils`

- Future extensions:
  - LLVAR / LLLVAR parser
  - Pretty ISO8583 message dumper

---

## Phase 3 — TPH Utilities (Transact)
- **tphutils**
  - **Length headers**
    - Add/remove 2-byte length headers
    - Add/remove 4-byte length headers
  - **TPDU header** handling (pack/unpack 5-byte TPDU)
  - **Framing helpers**: build message frame for TPH
  - **Mini TCP client**: send/receive framed messages (for testing)
  - CLI: input raw hex → output parsed frame (length + TPDU + body)

---

## Phase 4 — Crypto / Payments Utilities
- **cryptoutils**
  - Luhn check (PAN validation)
  - SHA256 / Base64 encode-decode
  - DES/3DES skeleton for MAC calculation

---

## Phase 5 — Fintech Utilities
- **finutils**
  - ISO4217 currency code lookup
  - Amount formatting with exponent
  - PAN masking (`411111******1111`)
  - ISO8583 datetime helpers (YYMMDDhhmmss, etc.)

---

## Phase 6 — Repo Polish
- `.gitignore` (done)
- LICENSE (MIT)
- CHANGELOG.md (start from v0.1.0)
- GitHub Actions CI (build + tests)
- README with examples for each tool
- Badges (Build ✅, Go version, License)

---

## 🚀 Milestones
- **v0.1.0** — numconvert polished
- **v0.2.0** — iso8583utils bitmap basics
- **v0.3.0** — tphutils length headers + TPDU helpers
- **v0.4.0** — cryptoutils (PAN + hashing)
- **v0.5.0+** — finutils & extra ISO8583 field parsers
