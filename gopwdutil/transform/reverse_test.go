package transform

import (
	"testing"
)

func ptr(b []byte) *[]byte { return &b }

func TestReverse_Nil(t *testing.T) {
	if err := Reverse(nil); err == nil {
		t.Error("expected error for nil input")
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("abcd"), []byte("dcba")},
		{[]byte("a"), []byte("a")},
		{[]byte(""), []byte("")},
	}
	for _, tt := range tests {
		p := ptr(tt.input)
		if err := Reverse(p); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(*p) != string(tt.want) {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, *p, tt.want)
		}
	}
}
