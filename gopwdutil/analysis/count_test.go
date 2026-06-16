package analysis

import "testing"

func ptr(b []byte) *[]byte { return &b }

func TestCountChar(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("hello"), 5},
		{[]byte("hello world"), 10},
		{[]byte(""), 0},
	}
	for _, tt := range tests {
		got, err := countChar(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countChar(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestCountChar_Nil(t *testing.T) {
	if _, err := countChar(nil); err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCountWord(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("hello world"), 2},
		{[]byte("one"), 1},
		{[]byte(""), 0},
		{[]byte("   "), 0},
	}
	for _, tt := range tests {
		got, err := countWord(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countWord(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestCountLetter(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("abc123"), 3},
		{[]byte("ABC"), 3},
		{[]byte("123!@#"), 0},
	}
	for _, tt := range tests {
		got, err := countLetter(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countLetter(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestCountUpper(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("Hello"), 1},
		{[]byte("HELLO"), 5},
		{[]byte("hello"), 0},
	}
	for _, tt := range tests {
		got, err := countUpper(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countUpper(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestCountLower(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("Hello"), 4},
		{[]byte("HELLO"), 0},
		{[]byte("hello"), 5},
	}
	for _, tt := range tests {
		got, err := countLower(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countLower(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestCountNumeric(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("abc123"), 3},
		{[]byte("abc"), 0},
		{[]byte("123"), 3},
	}
	for _, tt := range tests {
		got, err := countNumeric(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countNumeric(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestCountSpecial(t *testing.T) {
	tests := []struct{ input []byte; want int }{
		{[]byte("abc!@#"), 3},
		{[]byte("abc"), 0},
		{[]byte("!@#"), 3},
	}
	for _, tt := range tests {
		got, err := countSpecial(ptr(tt.input))
		if err != nil || got != tt.want {
			t.Errorf("countSpecial(%q) = %d, %v; want %d", tt.input, got, err, tt.want)
		}
	}
}

func TestGetCountUnit(t *testing.T) {
	tests := []struct{ count, choice int; want string }{
		{1, 1, "character"},
		{2, 1, "characters"},
		{0, 1, ""},
		{1, 2, "word"},
		{3, 2, "words"},
	}
	for _, tt := range tests {
		if got := getCountUnit(tt.count, tt.choice); got != tt.want {
			t.Errorf("getCountUnit(%d, %d) = %q, want %q", tt.count, tt.choice, got, tt.want)
		}
	}
}
