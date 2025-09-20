package numconvert

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput prompts the user for input via stdin.
// It supports numbers with the following prefixes:
//   - 0B for binary (e.g., 0B1010)
//   - 0X for hexadecimal (e.g., 0X1A)
//   - leading 0 for octal (e.g., 0755)
//   - no prefix → decimal (e.g., 1234)
//
// The input is normalized to uppercase and trimmed before being returned.
func ReadInput() string {
	fmt.Print("Enter a number (prefix with 0B for binary, 0X for hex): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	return strings.TrimSpace(input)
}

// DetectBase inspects the input string and returns its numeric base.
// Rules:
//   - "0B" prefix → base 2 (binary)
//   - "0X" prefix → base 16 (hexadecimal)
//   - leading "0" (with more than 1 digit) → base 8 (octal)
//   - otherwise → base 10 (decimal)
//
// If hex characters (A–F) are present, base 16 is returned.
func DetectBase(s string) int {
	hex := false

	if strings.HasPrefix(s, "0X") {
		return 16
	}
	if strings.HasPrefix(s, "0B") {
		return 2
	}
	if strings.HasPrefix(s, "0") && len(s) > 1 {
		for i := 1; i < len(s); i++ {
			if s[i] < '0' || s[i] > '7' {
				return 10
			}
		}
		return 8
	}

	for _, char := range s {
		if char >= 'A' && char <= 'F' {
			hex = true
		}
	}

	if hex {
		return 16
	}
	return 10
}

// ToDecimal converts a string s in the given base to its decimal (int64) value.
// It automatically strips "0X"/"0B" prefixes and uppercases the input.
// Returns an error if parsing fails (invalid character for the base).
func ToDecimal(s string, base int) (int64, error) {
	if strings.HasPrefix(s, "0X") || strings.HasPrefix(s, "0B") {
		s = s[2:]
		s = strings.ToUpper(s)
	}
	return strconv.ParseInt(s, base, 64)
}

// ToHex converts an int64 value to its uppercase hexadecimal representation,
// prefixed with "0X" (e.g., 26 → "0X1A").
func ToHex(n int64) string {
	return fmt.Sprintf("0X%X", n)
}

// ToBinary converts an int64 value to its binary representation,
// prefixed with "0B". It pads to 8 bits (e.g., 10 → "0B00001010").
func ToBinary(n int64) string {
	return fmt.Sprintf("0B%08b", n)
}

// ToOctal converts an int64 value to its octal representation,
// prefixed with "0" (e.g., 83 → "0123").
func ToOctal(n int64) string {
	return fmt.Sprintf("0%o", n)
}

// EncodeBCD encodes a string of decimal digits into BCD (Binary-Coded Decimal).
// Each byte stores two digits: high nibble = first digit, low nibble = second digit.
// Example: "1234" → []byte{0x12, 0x34}.
// Odd-length inputs are left-padded with a zero (e.g., "123" → "0123").
// Returns an error if the input contains non-digit characters.
func EncodeBCD(digits string) ([]byte, error) {
	if len(digits) == 0 {
		return nil, fmt.Errorf("input string is empty")
	}
	if len(digits)%2 != 0 {
		digits = "0" + digits
	}

	bcd := make([]byte, (len(digits)+1)/2)

	for i, char := range digits {
		if char < '0' || char > '9' {
			return nil, fmt.Errorf("invalid character: %c", char)
		}
		if i%2 == 0 {
			bcd[i/2] = byte((char - '0') << 4)
		} else {
			bcd[i/2] = bcd[i/2] | byte(char-'0') // use logical OR to set the lower nibble
		}
	}

	return bcd, nil
}

// DecodeBCD decodes a slice of BCD-encoded bytes into its digit string.
// Each byte is split into high and low nibbles, which must be valid digits (0–9).
// Example: []byte{0x12, 0x34} → "1234".
// Returns an error if any nibble is invalid (>9).
func DecodeBCD(bcd []byte) (string, error) {
	if len(bcd) == 0 {
		return "", fmt.Errorf("input byte slice is empty")
	}
	var digits strings.Builder
	for _, b := range bcd {
		highNibble := (b >> 4) & 0x0F
		lowNibble := b & 0x0F

		if highNibble > 9 || lowNibble > 9 {
			return "", fmt.Errorf("invalid BCD byte: %X", b)
		}

		digits.WriteByte('0' + highNibble)
		digits.WriteByte('0' + lowNibble)
	}

	return digits.String(), nil
}

// PrintResults prints the decimal, binary, hexadecimal, and octal
// representations of a given int64 value to stdout.
func PrintResults(n int64) {
	fmt.Println("Decimal:", n)
	fmt.Println("Binary:", ToBinary(n))
	fmt.Println("Hex:", ToHex(n))
	fmt.Println("Octal:", ToOctal(n))
}
