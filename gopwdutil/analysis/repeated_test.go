package analysis

import "testing"

func TestGetRepeated_Nil(t *testing.T) {
	if _, err := getRepeated(nil); err == nil {
		t.Error("expected error for nil input")
	}
}

func TestGetRepeated(t *testing.T) {
	tests := []struct {
		input    []byte
		wantKeys []byte
		wantNone bool
	}{
		{[]byte("aab"), []byte{'a'}, false},
		{[]byte("abcd"), nil, true},
		{[]byte("aabb"), []byte{'a', 'b'}, false},
	}
	for _, tt := range tests {
		got, err := getRepeated(ptr(tt.input))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if tt.wantNone && len(got) != 0 {
			t.Errorf("getRepeated(%q): expected no repeats, got %v", tt.input, got)
		}
		for _, k := range tt.wantKeys {
			if _, ok := got[k]; !ok {
				t.Errorf("getRepeated(%q): expected key %q in result", tt.input, k)
			}
		}
	}
}
