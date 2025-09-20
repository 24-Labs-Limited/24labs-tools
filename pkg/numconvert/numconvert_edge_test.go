package numconvert

import "testing"

func TestDetectBase_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{{
		"Empty string", "", 10},
		{"Weird prefix", "0Z123", 10},
		{"Hex without prefix", "ABCD", 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := DetectBase(tt.input); result != tt.expected {
				t.Errorf("DetectBase(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToDecimal_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		base    int
		wantErr bool
	}{
		{"Empty string", "", 10, true},
		{"Invalid binary", "0B102", 2, true},
		{"Invalid hex", "0X1G", 16, true},
		{"Too large", "9223372036854775808", 10, true}, // > MaxInt64
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToDecimal(tt.input, tt.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDecimal(%q, %d) error = %v, wantErr %v", tt.input, tt.base, err, tt.wantErr)
			}
		})
	}
}

func TestEncodeBCD_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
		wantErr  bool
	}{
		{"Empty string", "", nil, true},
		{"Odd length string", "123", []byte{0x01, 0x23}, false}, // padded to "0123"
		{"Invalid digit", "12A", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := EncodeBCD(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeBCD(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			// Only check result if no error expected
			if !tt.wantErr && string(result) != string(tt.expected) {
				t.Errorf("EncodeBCD(%q) = %X, expected %X", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDecodeBCD_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
		wantErr  bool
	}{
		{"Empty slice", []byte{}, "", true},
		{"Valid BCD", []byte{0x12, 0x34}, "1234", false},
		{"Invalid nibble", []byte{0x1F}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecodeBCD(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBCD(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			if result != tt.expected {
				t.Errorf("DecodeBCD(%v) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
