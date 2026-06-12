package tools

import (
	"strings"
)

func IsWhiteSpace(pwd string) bool {
	return len(strings.Fields(pwd)) == 0
}