package tools

import (
	"bufio"
	"errors"
	"os"
)

var Reader = bufio.NewReader(os.Stdin)

const MinPassword int = 8
const MaxPassword int = 72
const MinSalt int = 16
const MaxSalt int = 32
const MinPepper int = 16
const MaxPepper int = 32

type Password struct {
	Password []byte
	Salt []byte
	Pepper []byte
	HashedKey []byte
}

// Character sets used throughout the application
var LowerChar = []byte("abcdefghijklmnopqrstuvwxyz")
var UpperChar = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var NumericChar = []byte("0123456789")
var SpecialChar = []byte(" .,;:!?'\"()[]{}<>-_/\\|+*=%^@#$&~`")

var Errors = struct {
	NilError error
	EmptyError error
	WhiteSpace error
	LengthError	 error
	CaptureError error
	InvalidInput error
	InternalError error
} {
	errors.New("password is nil"),
	errors.New("empty string"),
	errors.New("whitespace only"),
	errors.New("invalid string length"),
	errors.New("capture error"),
	errors.New("invalid input error"),
	errors.New("internal error"),
}
