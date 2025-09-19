package numconvert

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() string {
	fmt.Print("Enter a number (prefix with 0B for binary, 0X for hex): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	return strings.TrimSpace(input)
}

func DetectBase(s string) int {
	hex := false

	if strings.HasPrefix(s, "0X") {
		return 16
	}
	if strings.HasPrefix(s, "0B") {
		return 2
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

func ToDecimal(s string, base int) (int64, error) {
	if strings.HasPrefix(s, "0X") || strings.HasPrefix(s, "0B") {
		s = s[2:]
		s = strings.ToUpper(s)
	}
	return strconv.ParseInt(s, base, 64)
}

func ToHex(n int64) string {
	return fmt.Sprintf("0X%X", n)
}

func ToBinary(n int64) string {
	return fmt.Sprintf("0B%08b", n)
}

func PrintResults(n int64) {
	fmt.Println("Decimal:", n)
	fmt.Println("Binary:", ToBinary(n))
	fmt.Println("Hex:", ToHex(n))
}
