package tools

import "testing"

func TestContainsByte(t *testing.T) {
	tests := []struct {
		b    byte
		ref  []byte
		want bool
	}{
		{'a', LowerChar, true},
		{'Z', UpperChar, true},
		{'5', NumericChar, true},
		{'!', SpecialChar, true},
		{'a', UpperChar, false},
		{'Z', LowerChar, false},
		{'a', []byte{}, false},
	}

	for _, tt := range tests {
		if got := ContainsByte(tt.b, tt.ref); got != tt.want {
			t.Errorf("ContainsByte(%q, ...) = %v, want %v", tt.b, got, tt.want)
		}
	}
}
