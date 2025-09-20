package numconvert

import (
	"reflect"
	"testing"
)

func TestDetectBase(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"0B1010", 2},
		{"0X1A", 16},
		{"0123", 8},
		{"1A2F", 16},
		{"123", 10},
		{"1010", 10},  // no prefix → decimal
		{"0Z123", 10}, // weird prefix → decimal
	}

	for _, tt := range tests {
		if got := DetectBase(tt.input); got != tt.expected {
			t.Errorf("DetectBase(%q) = %d; want %d", tt.input, got, tt.expected)
		}
	}
}

func TestToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		base     int
		expected int64
		wantErr  bool
	}{
		{"0B1010", 2, 10, false},
		{"0X1A", 16, 26, false},
		{"0123", 8, 83, false},
		{"123", 10, 123, false},
		{"12A", 10, 0, true}, // invalid for base 10
	}

	for _, tt := range tests {
		got, err := ToDecimal(tt.input, tt.base)
		if (err != nil) != tt.wantErr {
			t.Errorf("ToDecimal(%q, %d) error = %v, wantErr %v", tt.input, tt.base, err, tt.wantErr)
			continue
		}
		if got != tt.expected {
			t.Errorf("ToDecimal(%q, %d) = %d; want %d", tt.input, tt.base, got, tt.expected)
		}
	}
}

func TestToHex(t *testing.T) {
	if got := ToHex(26); got != "0X1A" {
		t.Errorf("ToHex(26) = %q; want 0X1A", got)
	}
}

func TestToBinary(t *testing.T) {
	if got := ToBinary(10); got != "0B00001010" {
		t.Errorf("ToBinary(10) = %q; want 0B00001010", got)
	}
}

func TestToOctal(t *testing.T) {
	if got := ToOctal(83); got != "0123" {
		t.Errorf("ToOctal(83) = %q; want 0123", got)
	}
}

func TestEncodeDecodeBCD(t *testing.T) {
	t.Run("Valid even length", func(t *testing.T) {
		bcd, err := EncodeBCD("1234")
		if err != nil {
			t.Fatalf("EncodeBCD error = %v", err)
		}
		expected := []byte{0x12, 0x34}
		if !reflect.DeepEqual(bcd, expected) {
			t.Errorf("EncodeBCD(1234) = %v; want %v", bcd, expected)
		}

		decoded, err := DecodeBCD(bcd)
		if err != nil {
			t.Fatalf("DecodeBCD error = %v", err)
		}
		if decoded != "1234" {
			t.Errorf("DecodeBCD(%v) = %q; want 1234", bcd, decoded)
		}
	})

	t.Run("Valid odd length", func(t *testing.T) {
		bcd, err := EncodeBCD("123")
		if err != nil {
			t.Fatalf("EncodeBCD error = %v", err)
		}
		expected := []byte{0x01, 0x23} // padded
		if !reflect.DeepEqual(bcd, expected) {
			t.Errorf("EncodeBCD(123) = %v; want %v", bcd, expected)
		}

		decoded, err := DecodeBCD(bcd)
		if err != nil {
			t.Fatalf("DecodeBCD error = %v", err)
		}
		if decoded != "0123" {
			t.Errorf("DecodeBCD(%v) = %q; want 0123", bcd, decoded)
		}
	})

	t.Run("Invalid char", func(t *testing.T) {
		if _, err := EncodeBCD("12A"); err == nil {
			t.Errorf("EncodeBCD(12A) expected error, got nil")
		}
	})

	t.Run("Invalid nibble", func(t *testing.T) {
		invalid := []byte{0x1A} // A = 10 not valid in BCD
		if _, err := DecodeBCD(invalid); err == nil {
			t.Errorf("DecodeBCD(%v) expected error, got nil", invalid)
		}
	})
}
