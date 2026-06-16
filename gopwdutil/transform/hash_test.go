package transform

import (
	"gopwdutil/tools"
	"testing"
)

func TestHash_Nil(t *testing.T) {
	if err := Hash(nil); err == nil {
		t.Error("expected error for nil input")
	}
}

func TestHash_ProducesNonEmptyKey(t *testing.T) {
	p := &tools.Password{
		Password: []byte("TestP@ssw0rd"),
		Salt:     []byte("somesalt1234567890"),
		Pepper:   []byte("somepepper123456789"),
	}
	if err := Hash(p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(p.HashedKey) == 0 {
		t.Error("expected non-empty hashed key")
	}
}

func TestHash_Deterministic(t *testing.T) {
	p1 := &tools.Password{Password: []byte("pass"), Salt: []byte("salt1234567890123456"), Pepper: []byte("pepper12345678901234")}
	p2 := &tools.Password{Password: []byte("pass"), Salt: []byte("salt1234567890123456"), Pepper: []byte("pepper12345678901234")}
	Hash(p1)
	Hash(p2)
	if string(p1.HashedKey) != string(p2.HashedKey) {
		t.Error("same input should produce same hash")
	}
}

func TestHash_DifferentSaltsDifferentHashes(t *testing.T) {
	p1 := &tools.Password{Password: []byte("pass"), Salt: []byte("salt1111111111111111")}
	p2 := &tools.Password{Password: []byte("pass"), Salt: []byte("salt2222222222222222")}
	Hash(p1)
	Hash(p2)
	if string(p1.HashedKey) == string(p2.HashedKey) {
		t.Error("different salts should produce different hashes")
	}
}
