# 📜 Changelog

## [Unreleased]

### Added

- `numconvert`: Base detection rules documented in `README.md` (now explicit and detailed).
- `numconvert`: New **safe octal detection** logic.  
  - Numbers starting with `0` are treated as **octal only if all digits are between `0–7`**.  
  - Otherwise, they fall back to **decimal**.  
  - Example:  
    - `0123` → octal  
    - `0Z123` → decimal  

### Fixed

- Tests updated to align with safe octal behavior (`0Z123` no longer misclassified as octal).

---

## [v0.1.0] - 2025-09-19

### Added

- **numconvert** package with support for:
  - Automatic base detection:
    - `0B` prefix → binary  
    - `0X` prefix → hexadecimal  
    - leading `0` → octal  
    - otherwise → decimal  
  - Conversions:
    - `ToDecimal` (string → int64)  
    - `ToHex` (int64 → hex string)  
    - `ToBinary` (int64 → binary string)  
    - `ToOctal` (int64 → octal string)  
  - **BCD utilities**:
    - `EncodeBCD` (string of digits → BCD bytes)  
    - `DecodeBCD` (BCD bytes → string of digits)  
  - CLI support via `cmd/numconvert` for quick conversions.
- Unit tests with ~79% coverage across core functions.
- Example usage in README.
