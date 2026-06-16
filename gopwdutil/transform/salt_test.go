package transform

import (
	"gopwdutil/tools"
	"testing"
)

func TestSalt_Nil(t *testing.T) {
	if err := Salt(nil); err == nil {
		t.Error("expected error for nil input")
	}
}

func TestSalt_LengthInRange(t *testing.T) {
	salt := []byte{}
	if err := Salt(&salt); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(salt) < tools.MinSalt || len(salt) >= tools.MaxSalt {
		t.Errorf("salt length %d not in [%d, %d)", len(salt), tools.MinSalt, tools.MaxSalt)
	}
}

func TestSalt_Randomness(t *testing.T) {
	salt1, salt2 := []byte{}, []byte{}
	Salt(&salt1)
	Salt(&salt2)
	if string(salt1) == string(salt2) {
		t.Error("two generated salts should not be equal (extremely unlikely)")
	}
}
